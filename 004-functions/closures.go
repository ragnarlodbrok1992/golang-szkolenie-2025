package main

import "fmt"

// outerFunction returns a closure that captures the variable 'x'
func outerFunction(x int) func() int {
    // innerFunction is a closure that captures 'x'
    innerFunction := func() int {
        x++ // x is captured and modified
        return x
    }
    return innerFunction
}

func PrintClosuresInfo() {
	fmt.Println("Hello from Closures!")

  // Create a closure by calling outerFunction with an initial value of 10
  closure := outerFunction(10)

  // Call the closure multiple times to see the captured variable change
  fmt.Println(closure()) // Output: 11
  fmt.Println(closure()) // Output: 12
  fmt.Println(closure()) // Output: 13

  // Create another closure with a different initial value
  anotherClosure := outerFunction(20)
  fmt.Println(anotherClosure()) // Output: 21
  fmt.Println(anotherClosure()) // Output: 22
}
