package main

import (
	"fmt"
	"errors"
)

// Define a struct to group data
type Person struct {
	Name string
	Age int
}

// Defining a functions
func greet(person Person) string {
	return fmt.Sprintf("Hello, %s! You are %d years old.", person.Name, person.Age)
}

// Function with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}


func main() {
	fmt.Println("Hello, world!")

	// Variables and constants
	const greeting = "Hello, World from a constant!"
	var name string = "Alice"
	age := 30 // Short declaration

	fmt.Println(greeting)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Using structs
	person := Person{Name: "Bob", Age: 25}
	fmt.Println(greet(person))

	// Using arrays and slices
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// Loops of course
	for id, num := range nums {
		fmt.Printf("In slot %d we have a number --> %d\n", id, num)
	}

	// Using maps
	ages := map [string] int {
		"Alice": 30,
		"Zenek": 13,
		"Charlie": 52,
		"Bob": 25,
	}
	fmt.Println("Ages:", ages)

	// Using pointers
	ptr := &name
	fmt.Printf("Name via pointer: %s\n", *ptr)

	// Error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// Conditionals
	if age > 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult - terminating...")
	}
}
