#!/usr/bin/env bash
set -euo pipefail

# The directory where we store our templates:
TEMPLATES_DIR=".templar"

# 1. Ensure the .templar directory exists
if [ ! -d "$TEMPLATES_DIR" ]; then
  echo "ERROR: No $TEMPLATES_DIR directory found in $(pwd)."
  echo "Please create one or run this script from the correct location."
  exit 1
fi

echo "Rendering templates from '$TEMPLATES_DIR' ..."

# 2. Recursively find all regular files within .templar
find "$TEMPLATES_DIR" -type f -name "*.tmpl" | while read -r template_file; do
  # 3. Compute the relative path without the .templar prefix
  #    e.g., ".templar/etc/nginx/nginx.config" -> "etc/nginx/nginx.config"
  relative_path="${template_file#"$TEMPLATES_DIR"/}"

  # 4. Create the subdirectories in the current directory (if they don’t exist)
  mkdir -p "$(dirname "$relative_path")"

  # 5. Render the template via stdin -> templar -> output file
  templar < "$template_file" > "${relative_path%.tmpl}"

  echo "Rendered '$template_file' -> '${relative_path%.tmpl}'"
done

echo "All templates have been rendered successfully."

