package main

import (
	"fmt"
	"sync"
	"time"
)

// Example functions
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Synchronization
// link --> https://pkg.go.dev/sync#WaitGroup
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Wait group is just a counter for workers - starting worker increments it, calling Done() decrements it

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}


func main() {
	fmt.Println("Hello from concurrency!")

	// Normal synchrounous run
	say("hello...")
	say("world!")

	// Asynchronous goroutine run
	go say("hello...")
	say("world?")

	// waiting for asynchrounous functions
	var wait_group sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wait_group.Add(1)
		go worker(i, &wait_group)
	}

	fmt.Println("Before call to Wait() in WaitGroup...")
	wait_group.Wait()
	fmt.Println("After call to Wait() in WaitGroup!")

	// Goroutines
	PrintGoroutinesLesson()

	// Channels
	PrintChannelsLesson()

	// Channels - chapter 2
	PrintChannelsSecondLesson()
}
