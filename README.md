# MiniTMPL

MiniTMPL is a simple, yet powerful tool designed to render text templates using environment variables. It reads a template from `stdin` and outputs the rendered version to `stdout`, making it an excellent tool for generating configuration files, messages, or any text-based content that requires dynamic data insertion.

## Features

- Reads templates from standard input (`stdin`).
- Renders templates using environment variables.
- Supports basic template functions provided by Go's `text/template` package.

## Getting Started

To use EnvTemplater, clone this repository or download the latest release to your local machine.

### Prerequisites

Ensure you have Go installed on your system. EnvTemplater requires Go 1.14 or higher.


### Building form source

First, clone the repository:

```bash
git clone https://git.behzadan.ir/p/minitmpl.git
cd minitmpl
```

Then, build the program:

```bash
go build -o minitmpl
```

## Usage

To use EnvTemplater, simply pipe a template into the program and it will output the rendered version:

```bash
echo "Hello, {{.USER}}" | ./envtemplater
```

This will replace `{{.USER}}` with the value of the `USER` environment variable.

## License

This project is licensed under a custom license - see the [LICENSE](LICENSE) file for details.

