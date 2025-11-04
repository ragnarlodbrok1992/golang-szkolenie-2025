package main

import "fmt"

// Basic function
func add(a int, b int) int {
    return a + b
}

// Multiple return values
func divmod(a, b int) (int, int, error) {
    if b == 0 {
        return 0, 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, a % b, nil
}

// Named return values
func split(total int) (x, y int) {
    x = total * 4 / 10
    y = total - x
    return // "naked" return
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Method on a struct
type Rectangle struct {
    width, height int
}

func (r Rectangle) area() int {
    return r.width * r.height
}

// Function with defer
func process() {
    defer fmt.Println("Deferred statement")
    fmt.Println("Processing")
}

func main() {
	fmt.Println("Hello from functions!")
  // Basic function
  fmt.Println("Add:", add(5, 3))

  // Multiple return values
  quotient, remainder, err := divmod(10, 3)
  if err != nil {
      fmt.Println("Error:", err)
  } else {
      fmt.Println("Quotient:", quotient, "Remainder:", remainder)
  }

  // Named return values
  a, b := split(10)
  fmt.Println("Split:", a, b)

  // Variadic function
  fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

  // Anonymous function
  square := func(x int) int {
      return x * x
  }
  fmt.Println("Square:", square(5))

  // Method on a struct
  rect := Rectangle{width: 10, height: 5}
  fmt.Println("Area:", rect.area())

  // Defer statement
  process()

	// Recursion
	PrintRecursionInfo()

	// Closures
	PrintClosuresInfo()
}
