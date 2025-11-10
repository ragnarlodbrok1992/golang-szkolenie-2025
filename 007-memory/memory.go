package main

import "fmt"

// Return a pointer to heap-allocated memory
func returnPointer() *int {
    value := 42
    // The variable 'value' is allocated on the heap due to heap promotion.
    // This is because the pointer to 'value' escapes the local scope of the function.
    return &value
}

// Demonstrate memory management with pointers
func pointerExample() {
    // Using pointers
    var ptr *int
    value := 42
    ptr = &value
    fmt.Println("Value via pointer:", *ptr)
		fmt.Println("Pointer is:", ptr)
}

// Demonstrate memory allocation with 'new'
func newExample() {
    // Using 'new' to allocate memory
    ptr := new(int)
    *ptr = 100
    fmt.Println("Value via new pointer:", *ptr)
}

// Demonstrate memory allocation with 'make' for slices
func sliceExample() {
    // Using 'make' to create a slice
    slice := make([]int, 5, 10) // length 5, capacity 10
    slice[0] = 1
    slice[1] = 2
    slice[2] = 3
    fmt.Println("Slice:", slice)
}

// Demonstrate memory allocation with 'make' for maps
func mapExample() {
    // Using 'make' to create a map
    m := make(map[string]int)
    m["key1"] = 10
    m["key2"] = 20
    fmt.Println("Map:", m)
}

// Demonstrate passing memory to functions
func modifySlice(s []int) {
    s[0] = 100
}

// Demonstrate removing values from slices and maps
func removeFromSlice() {
    slice := make([]int, 3)
    slice[0] = 1
    slice[1] = 2
    slice[2] = 3
    fmt.Println("Slice before removal:", slice)

    // Remove the first element
    slice = slice[1:]
    fmt.Println("Slice after removal:", slice)
}

func removeFromMap() {
    m := make(map[string]int)
    m["key1"] = 10
    m["key2"] = 20
    fmt.Println("Map before removal:", m)

    // Remove a key from the map
    delete(m, "key1")
    fmt.Println("Map after removal:", m)
}

func main() {
    // Example with pointers
    pointerExample()

    // Example with 'new'
    newExample()

    // Example with slices
    sliceExample()

    // Example with maps
    mapExample()

    // Example of passing memory to functions
    slice := make([]int, 3)
    slice[0] = 1
    slice[1] = 2
    slice[2] = 3
    fmt.Println("Slice before modification:", slice)
    modifySlice(slice)
    fmt.Println("Slice after modification:", slice)

    // Example of returning a pointer to heap-allocated memory
    ptr := returnPointer()
    fmt.Println("Value via returned pointer:", *ptr)

    // Example of removing values from slices and maps
    removeFromSlice()
    removeFromMap()
}
