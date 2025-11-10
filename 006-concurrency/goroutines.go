package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task(id int) {
	rand.Seed(time.Now().UnixNano())
	duration := time.Duration(rand.Intn(3) + 1) * time.Second

	fmt.Printf("Task %d started, will take %v to complete\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Task %d completed\n", id)
}

func PrintGoroutinesLesson() {
	fmt.Println("Hello from Goroutines!")

	const numTasks = 5

	for i := 1; i <= numTasks; i++ {
		go task(i)
	}

	time.Sleep(5 * time.Second) // Comment out to see what happens
	fmt.Println("All tasks completed (eventually)")
}
