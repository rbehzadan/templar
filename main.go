package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"templar/functions"

	"dario.cat/mergo"
	"github.com/Masterminds/sprig/v3"
)

//go:embed VERSION
var versionFile embed.FS

func printHelp() {
	helpText := fmt.Sprintf(`
Usage: %s [OPTIONS]

This program reads a text template from stdin and uses environment variables to render it.

Options:
  -v      Prints the version of the program.
  -h      Prints this help message.

Example:
  echo "Hello, {{.USER}}" | %s
`, os.Args[0], os.Args[0])
	fmt.Print(helpText)
}

func main() {
	// Define and parse flags
	versionFlag := flag.Bool("v", false, "Prints the version of the program")
	helpFlag := flag.Bool("h", false, "Prints the help message")
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		version, err := versionFile.ReadFile("VERSION")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading version: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Version:", strings.TrimSpace(string(version)))
		os.Exit(0)
	}

	// Handle help flag
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// Define custom functions
	customFunctions := sprig.FuncMap()
	err := mergo.Merge(&customFunctions, functions.FuncMap())
	if err != nil {
		fmt.Printf("Error merging custom functions: %v\n", err)
		os.Exit(1)
	}

	// Create a new template.
	tmpl, err := template.New("template").Funcs(customFunctions).Parse(readInput())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		os.Exit(1)
	}

	// Create a map to hold environment variables.
	envVars := make(map[string]string)

	// Populate the map with the current environment variables.
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}

	// Execute the template with the environment variables map.
	if err := tmpl.Execute(os.Stdout, envVars); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
		os.Exit(1)
	}
}

// readInput reads from stdin until EOF and returns the input as a string.
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	return input
}
