package main

import "fmt"

func main() {
	fmt.Println("Hello from types --> 002-types and syntax.")
  // Boolean type
  var boolVar bool = true
  fmt.Println("Boolean:", boolVar)

  // String type
  var stringVar string = "Hello, World!"
  fmt.Println("String:", stringVar)

  // Integer types
  var intVar int = 42
  var int8Var int8 = 10
  var int16Var int16 = 20
  var int32Var int32 = 30
  var int64Var int64 = 40
  fmt.Println("Int:", intVar)
  fmt.Println("Int8:", int8Var)
  fmt.Println("Int16:", int16Var)
  fmt.Println("Int32:", int32Var)
  fmt.Println("Int64:", int64Var)

  // Unsigned integer types
  var uintVar uint = 52
  var uint8Var uint8 = 11
  var uint16Var uint16 = 21
  var uint32Var uint32 = 31
  var uint64Var uint64 = 41
  fmt.Println("Uint:", uintVar)
  fmt.Println("Uint8:", uint8Var)
  fmt.Println("Uint16:", uint16Var)
  fmt.Println("Uint32:", uint32Var)
  fmt.Println("Uint64:", uint64Var)

  // Float types
  var float32Var float32 = 3.14
  var float64Var float64 = 3.1415926535
  fmt.Println("Float32:", float32Var)
  fmt.Println("Float64:", float64Var)

  // Complex types
  var complex64Var complex64 = 1 + 2i
  var complex128Var complex128 = 1 + 3i
  fmt.Println("Complex64:", complex64Var)
  fmt.Println("Complex128:", complex128Var)

  // Byte and Rune types
  var byteVar byte = 65
  var runeVar rune = 'A'
  fmt.Println("Byte:", byteVar)
  fmt.Println("Rune:", runeVar)

  // Constant declaration
  const constVar int = 100
  fmt.Println("Constant:", constVar)

	// Importing local functions
	PrintStructsInfo()
}
