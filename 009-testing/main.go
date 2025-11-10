package main

import "fmt"

func main() {
	fmt.Println("Hello from testing examples, I am Your main function!")

	// How to run tests
	// go test . // dot is for current directory
	// go test -v .
	// go test -cover .
	// go test -bench .
	// go test -bench . -benchmem
}
