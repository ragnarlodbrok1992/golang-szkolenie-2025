package main

import "fmt"

func PrintControlFlow() {
  // If-else control flow
  x := 10
  if x > 10 {
      fmt.Println("x is greater than 10")
  } else if x == 10 {
      fmt.Println("x is equal to 10")
  } else {
      fmt.Println("x is less than 10")
  }

  // For loop control flow
  for i := 0; i < 5; i++ {
      fmt.Printf("Current iteration: %d\n", i)
  }

  // Switch control flow
  day := "Tuesday"
  switch day {
  case "Monday":
      fmt.Println("Today is Monday")
  case "Tuesday":
      fmt.Println("Today is Tuesday")
  case "Wednesday":
      fmt.Println("Today is Wednesday")
  default:
      fmt.Println("Today is another day")
  }
}
