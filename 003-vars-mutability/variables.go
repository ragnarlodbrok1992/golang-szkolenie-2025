package main

import "fmt"

// Function to modify a variable passed by value
func modifyByValue(x int) {
    x = 20
    fmt.Println("Inside modifyByValue, x is:", x)
}

// Function to modify a variable passed by reference (using a pointer)
func modifyByReference(x *int) {
    *x = 20
    fmt.Println("Inside modifyByReference, x is:", *x)
}

func main() {
	fmt.Println("Variables and mutability")
  // Mutable variable
  var x int = 10
  fmt.Println("Initial value of x:", x)
  x = 20
  fmt.Println("New value of x:", x)

  // Immutable variable
  const a int = 30
  fmt.Println("Value of y:", a)

  // Uncommenting the following line will cause a compilation error
  // y = 40

  // Example with pass-by-value
  x = 10
  fmt.Println("Before calling modifyByValue, x is:", x)
  modifyByValue(x)
  fmt.Println("After calling modifyByValue, x is:", x)

  // Example with pass-by-reference (using a pointer)
  var y = 10
  fmt.Println("Before calling modifyByReference, y is:", y)
  modifyByReference(&y)
  fmt.Println("After calling modifyByReference, y is:", y)

	// Control flow
	PrintControlFlow()
}
