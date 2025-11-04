package main

import "fmt"

// Define a basic struct for Address
type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

// Define a struct for Person with various types of fields
type Person struct {
    Name      string
    Age       int
    IsStudent bool
    Address   // Embedded struct
    Courses   []string // Slice of strings
    Grades    map[string]int // Map with string keys and int values
    CalculateGPA func([]int) float64 // Function field
}

// Method on the Person struct
func (p Person) PrintInfo() {
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("Is Student:", p.IsStudent)
    fmt.Println("Address:", p.Address)
    fmt.Println("Courses:", p.Courses)
    fmt.Println("Grades:", p.Grades)
}

// Functions available in other codefiles should start from capital letter
func PrintStructsInfo() {
	fmt.Println("Hello from Structs.go!")

	// Some code about structures
  // Create an instance of the Person struct
  person1 := Person{
      Name:      "Alice",
      Age:       25,
      IsStudent: true,
      Address: Address{
          Street:  "123 Main St",
          City:    "Anytown",
          State:   "CA",
          ZipCode: "12345",
      },
      Courses: []string{"Math", "Science", "History"},
      Grades: map[string]int{
          "Math":    90,
          "Science": 85,
          "History": 95,
      },
      CalculateGPA: func(grades []int) float64 {
          sum := 0
          for _, grade := range grades {
              sum += grade
          }
          return float64(sum) / float64(len(grades))
      },
  }

  // Access and modify the fields of the struct
  fmt.Println("Name:", person1.Name)
  fmt.Println("Age:", person1.Age)

  // Call the method on the struct instance
  person1.PrintInfo()

  // Use the function field
  grades := []int{90, 85, 95}
  gpa := person1.CalculateGPA(grades)
  fmt.Println("GPA:", gpa)
}
