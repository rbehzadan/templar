#!/usr/bin/env bash
set -euo pipefail

REPO_OWNER="rbehzadan"
REPO_NAME="templar"
BINARY="templar"
GIT_SERVER="github.com"
INSTALL_DIR="/usr/local/bin"

# 1. Check OS
OS="$(uname -s)"
if [[ "$OS" != "Linux" ]]; then
  echo "Error: only Linux is supported by this installer script."
  exit 1
fi

# 2. Detect architecture
ARCH_DETECTED="$(uname -m)"
case "$ARCH_DETECTED" in
  x86_64)
    ARCH="amd64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  armv7l|armv7)
    ARCH="armv7"
    ;;
  *)
    echo "Error: unsupported architecture '$ARCH_DETECTED'"
    exit 1
    ;;
esac

# 3. Check for required tools (curl, jq, tar)
for cmd in curl jq tar; do
  if ! command -v "$cmd" >/dev/null; then
    echo "Error: '$cmd' is required but not found in PATH."
    exit 1
  fi
done

# 4. Get the latest version tag (strip leading 'v' if present)
VERSION="$(curl -fsSL "https://api.${GIT_SERVER}/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest" \
  | jq -r ".tag_name" \
  | sed 's/^v//')"

if [[ -z "$VERSION" || "$VERSION" == "null" ]]; then
  echo "Error: could not detect the latest version from GitHub."
  exit 1
fi

# 5. Construct package name and download
PACKAGE="${BINARY}_${VERSION}_linux_${ARCH}.tar.gz"
REPO_URL="https://${GIT_SERVER}/${REPO_OWNER}/${REPO_NAME}"

TEMPDIR="$(mktemp -d)"
pushd "$TEMPDIR" >/dev/null

echo "Downloading $PACKAGE from $REPO_URL/releases/download/v${VERSION}/"
curl -fsSLO "${REPO_URL}/releases/download/v${VERSION}/${PACKAGE}"

# 6. Extract tar file
echo "Extracting $PACKAGE..."
tar xf "${PACKAGE}"

# 7. Install templar and render.sh
echo "Installing $BINARY to $INSTALL_DIR..."
sudo install "$BINARY" "$INSTALL_DIR"

if [[ -f "render.sh" ]]; then
  echo "Installing render.sh to $INSTALL_DIR..."
  sudo install "render.sh" "$INSTALL_DIR"
fi

popd >/dev/null
rm -rf "$TEMPDIR"

# 8. Print installed version
echo ""
echo "Installation complete."
echo "Templar version: $($BINARY -v || true)"

