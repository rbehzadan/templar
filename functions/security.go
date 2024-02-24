package functions

import (
	"crypto/rand"
	"fmt"
)

// genpw generates a random password of the given length.
// It uses a predefined set of characters to ensure the password is readable.
func genpw(length int) (string, error) {
	// Define a set of characters to use in the password.
	// You can adjust the character set to meet your password policy requirements.
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

	b := make([]byte, length) // Create a slice of bytes to store the password characters
	_, err := rand.Read(b)    // Fill the slice with random bytes
	if err != nil {
		return "", fmt.Errorf("error generating random bytes: %v", err)
	}

	password := make([]byte, length)
	for i, byteVal := range b {
		password[i] = charset[byteVal%byte(len(charset))] // Map each byte to a character in the charset
	}

	return string(password), nil // Convert the password slice to a string and return it
}

func GenPassword(length int) string {
	pw, err := genpw(length)
	if err != nil {
		return "7HQO0LfNL;={i,4@0Tz/]2#EC"
	}
	return pw
}
