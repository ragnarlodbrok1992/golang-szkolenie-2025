package main

import (
	"fmt"
	// "math/rand"
	// "sync"
	// "time"
)

func PrintChannelsLesson() {
	fmt.Println("Hello from Channels!")

	// Unbuffered channel example
	unbuffered_channel := make(chan int)

	go func() {
		fmt.Println("Sending value to unbuffered channel")
		unbuffered_channel <- 42
	}()

	value := <-unbuffered_channel
	fmt.Println("Received value from unbuffered channel:", value)

	// Buffered channel example
	buffered_channel := make(chan int, 3)
	fmt.Println("Sending values to buffered channel")
	buffered_channel <- 1
	buffered_channel <- 2
	buffered_channel <- 3

	fmt.Println("Receiving values from buffered channel.")
	for i := 0; i < 3; i++ {
		fmt.Println("Buffered channel value --> ", <-buffered_channel)
	}

	// Closing the channel
	closed_channel := make(chan string)
	go func() {
		// https://pkg.go.dev/builtin#close
		defer close(closed_channel)

		for i := 0; i < 3; i++ {
			closed_channel <- fmt.Sprintf("Message number %d", i + 1)
		}
	}()

	for i := 0; i < 3; i++ {
		msg, ok := <-closed_channel
		if ok {
			fmt.Println(msg)
		} else {
			fmt.Println("Channel is closed!")
		}
	}

	// Pushing value into channel after closing - careful!
	// closed_channel <- "Wait for me!"

	// Range iterator on channel
	some_channel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		some_channel <- i
	}

	close(some_channel) // Comment to see error
	for msg := range some_channel {
		fmt.Println("Value in channel is: ", msg)
	}
}
