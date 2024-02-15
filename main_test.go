package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestTemplateRendering tests rendering a template using an environment variable.
func TestTemplateRendering(t *testing.T) {
	// Backup the original stdin and stdout
	origStdin := os.Stdin
	origStdout := os.Stdout

	// Create a pipe to simulate stdin
	tempStdin, tempStdinFile, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe for stdin: %v", err)
	}
	os.Stdin = tempStdin

	// Create a pipe to capture stdout
	tempStdoutRead, tempStdoutWrite, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe for stdout: %v", err)
	}
	os.Stdout = tempStdoutWrite

	// Set environment variable for the test
	os.Setenv("USER", "testuser")
	defer func() {
		// Cleanup: Restore stdin, stdout and unset the environment variable
		os.Stdin = origStdin
		os.Stdout = origStdout
		os.Unsetenv("USER")
		tempStdoutWrite.Close()
	}()

	// Write the template to the tempStdinFile, simulating input
	input := "Hello, {{.USER}}!"
	go func() {
		tempStdinFile.Write([]byte(input))
		tempStdinFile.Close()
	}()

	// Run the program (assuming main does the reading, processing, and printing)
	main()

	// Close write end of stdout pipe to finish reading from it
	tempStdoutWrite.Close()

	// Read the output from the captured stdout
	output, err := io.ReadAll(tempStdoutRead)
	if err != nil {
		t.Fatalf("Failed to read captured stdout: %v", err)
	}

	// Verify the output
	expectedOutput := "Hello, testuser!"
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Expected output to contain %q, got %q instead", expectedOutput, output)
	}
}
