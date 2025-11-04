package main

import "fmt"

// Factorial function using recursion
func factorial(n int) int {
  if n == 0 {
      return 1
  }
  return n * factorial(n-1)
}

// Fibonacci function using recursion
func fibonacci(n int) int {
  if n <= 1 {
      return n
  }
  return fibonacci(n-1) + fibonacci(n-2)
}

// GCD function using recursion and the Euclidean algorithm
func gcd(a, b int) int {
  if b == 0 {
      return a
  }
  return gcd(b, a%b)
}

func PrintRecursionInfo() {
	fmt.Println("Hello from Recursion!")
  // Example usage of factorial
  num := 5
  result := factorial(num)
  fmt.Printf("Factorial of %d is %d\n", num, result)

  // Example usage of Fibonacci
  fmt.Print("Fibonacci sequence: ")
  for i := 0; i < 10; i++ {
      fmt.Printf("%d ", fibonacci(i))
  }
  fmt.Println()

  // Example usage of GCD
  a, b := 48, 18
  gcdResult := gcd(a, b)
  fmt.Printf("GCD of %d and %d is %d\n", a, b, gcdResult)
}

