package main

import (
	"fmt"
	"os"
	"strconv"
)

// CustomError is a custom error type that includes a message and a code.
type CustomError struct {
	Message string
	Code    int
}

// Implement the Error method for the CustomError type.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Function to read a file and return its content as a string.
func readFile(filename string) (string, error) {
	// Attempt to open the file.
	file, err := os.Open(filename)
	if err != nil {
		// Return the error if the file cannot be opened.
		return "", err
	}
	defer file.Close() // Ensure the file is closed when the function returns.

	// Read the file content.
	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		// Return the error if there is an issue reading the file.
		return "", err
	}

	// Return the file content as a string.
	return string(data), nil
}

// Function to parse a number from a string.
func parseNumber(s string) (int, error) {
	// Attempt to convert the string to an integer.
	num, err := strconv.Atoi(s)
	if err != nil {
		// Return a custom error if the conversion fails.
		return 0, &CustomError{Message: "Invalid number format", Code: 400}
	}
	return num, nil
}

// Function that demonstrates panic and recover.
func handlePanic() {
	// Use defer to ensure the recover function is called if a panic occurs.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate a panic.
	panic("Something went terribly wrong!")
}

// Main function to demonstrate error handling.
func main() {
	// Message that should be seen at the end of the function to ensure everythign went 'smoothly'
	defer fmt.Println("All went well, goodbye...")

	// Example of basic error handling.
	content, err := readFile("config.txt")
	if err != nil {
		// Handle the error by printing it.
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("File content:", content)
	}

	// Example of custom error handling.
	num, err := parseNumber("123a")
	if err != nil {
		// Check if the error is of type CustomError.
		if customErr, ok := err.(*CustomError); ok {
			fmt.Printf("Custom Error: %v\n", customErr)
		} else {
			fmt.Println("Error parsing number:", err)
		}
	}
	fmt.Println("Parsed number:", num)

	// Example of panic and recover.
	handlePanic()

	// Example of error wrapping (Go 1.13+).
	// Suppose we have a function that can return an error.
	// We can wrap the error with additional context.
	_, err = readFile("nonexistent.txt")
	if err != nil {
		// Wrap the error with additional context.
		wrappedErr := fmt.Errorf("failed to read file: %w", err)
		fmt.Println("Wrapped error:", wrappedErr)
	}
}
