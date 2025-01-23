# Templar

**Templar** is a simple, yet powerful CLI tool for rendering Go text templates, enriched by environment variables. It reads a template from standard input (`stdin`) and outputs the rendered result to standard output (`stdout`). This makes it perfect for dynamically generating configuration files, messages, or any text-based content that needs environment-aware templating.

## Features

- **Simple Input/Output**: Pipe any text-based template into Templar, and receive fully rendered output.
- **Environment Variables**: Inserts environment variables into your templates (`{{.VAR}}`).
- **Built-in Template Functions**:
  - All functions from [Sprig](https://github.com/Masterminds/sprig).
  - [Additional custom functions](docs/custom_functions.md) for string manipulation, domain extraction, and more.

## Getting Started

### Installation

You can either:

1. **Download the Latest Release** (Recommended)  
   Check the [Releases](https://github.com/rbehzadan/templar/releases) page for precompiled binaries.

2. **Build from Source**  
   Requires [Go 1.14+](https://go.dev/dl/).

   ```bash
   git clone https://github.com/rbehzadan/templar.git
   cd templar
   go build -o templar
   ```

### Usage

Templar reads from `stdin` and writes to `stdout`. Here’s a simple example:

```bash
echo "Hello, {{.USER}}" | ./templar
```

- The template `Hello, {{.USER}}` is processed.
- `USER` is replaced by the value of the `USER` environment variable on your system.

**Using Custom Functions**  
Templar includes functions from [Sprig](https://github.com/Masterminds/sprig) plus [custom ones](CUSTOM_FUNCTIONS.md) like `title`, `split`, `genpw`, and more. For example:

```bash
echo "Password: {{genpw 16}}" | ./templar
```

This generates a new 16-character random password each time you run it.

## Batch Rendering with `render.sh`

In addition to the basic `templar` usage, this repository includes a helper script named [`render.sh`](./render.sh). It automates the process of rendering **multiple `.tmpl` files**:

1. **Create a `.templar/` directory** in the root of your project.
2. **Place your template files** (ending in `.tmpl`) inside that directory. For example:
   ```
   .templar/
   └── etc/
       └── nginx/
           └── nginx.conf.tmpl
   ```
3. **Run `./render.sh`** from the project root:
   ```bash
   chmod +x render.sh
   ./render.sh
   ```
   - The script recursively searches `.templar/` for any `.tmpl` files.
   - For each file, it removes the `.tmpl` extension and writes the rendered output to the corresponding location in your current directory.  
   - In the above example, `.templar/etc/nginx/nginx.conf.tmpl` is rendered to `etc/nginx/nginx.conf`.

> **Note:** Make sure you have `templar` either in your `$PATH` or in the same directory so that `render.sh` can call it.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

