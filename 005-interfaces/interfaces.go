package main

import "fmt"

// Define an interface named Shape with a method Area() that returns a float64
type Shape interface {
	Area() float64
}

// Define a struct named Circle
type Circle struct {
	Radius float64
}

// Implement the Area() method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Define a struct named Rectangle
type Rectangle struct {
	Width, Height float64
}

// Implement the Area() method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to print the area of any Shape
func PrintArea(s Shape) {
  fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	fmt.Println("Hello from interfaces!")

  // Create instances of Circle and Rectangle
  circle := Circle{Radius: 5}
  rectangle := Rectangle{Width: 10, Height: 5}

  // Call PrintArea with Circle and Rectangle instances
  PrintArea(circle)
  PrintArea(rectangle)
}
