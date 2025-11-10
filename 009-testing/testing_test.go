package main

// https://pkg.go.dev/testing <-- check this out

import (
	"testing"
	// The testing package provides support for automated testing of Go packages.
)

// IntMin is a function that returns the minimum of two integers.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestIntMin is a test function for IntMin.
func TestIntMinBasic(t *testing.T) {
	// A test function must start with the word 'Test' and take a pointer to testing.T as its only argument.
	// This is a basic test case.
	result := IntMin(1, 2)
	if result != 1 {
		// t.Error is used to indicate that the test has failed but continues execution.
		t.Error("Expected 1, got", result)
	}
}

// Wrong result // FIXME: fix before checking benches
func TestIntMinWrong(t *testing.T) {
	// A test function must start with the word 'Test' and take a pointer to testing.T as its only argument.
	// This is a basic test case.
	result := IntMin(1, 2)
	if result != 2 {
		// t.Error is used to indicate that the test has failed but continues execution.
		t.Error("Expected 2, got", result)
	}
}

// TestIntMinTableDriven demonstrates table-driven testing.
func TestIntMinTableDriven(t *testing.T) {
	// Table-driven tests allow you to test multiple scenarios with a single test function.
	// Each test case is defined as a struct.
	tests := []struct {
		name     string // Name of the test case
		a, b     int    // Input values
		expected int    // Expected result
	}{
		{"case1", 1, 2, 1},
		{"case2", 0, -1, -1},
		{"case3", -10, 10, -10},
	}

	for _, test := range tests {
		// t.Run allows running each test case as a subtest.
		t.Run(test.name, func(t *testing.T) {
			result := IntMin(test.a, test.b)
			if result != test.expected {
				// t.Errorf is used to indicate a failed test with a formatted message.
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

// TestIntMinWithSetup demonstrates setup and teardown.
func TestIntMinWithSetup(t *testing.T) {
	// Setup code can be placed here.
	// For example, initializing resources.
	setup := func() {
		// Setup logic goes here.
	}
	teardown := func() {
		// Teardown logic goes here.
	}

	setup()
	defer teardown()

	// Test logic goes here.
	result := IntMin(5, 3)
	if result != 3 {
		t.Error("Expected 3, got", result)
	}
}

// BenchmarkIntMin is a benchmark function for IntMin.
func BenchmarkIntMin(b *testing.B) {
	// Benchmark functions must start with the word 'Benchmark' and take a pointer to testing.B as its only argument.
	// Benchmark functions measure the performance of a function.
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
