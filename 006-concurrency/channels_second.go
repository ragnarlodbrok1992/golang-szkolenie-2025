package main

import (
	"fmt"
	"time"
)

func PrintChannelsSecondLesson() {
	fmt.Println("Channels - second lesson!")

	// Select statement with channels
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel_1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel_2 <- "from chnanel 2"
	}()

	// We might want to wait for only one value from one channel
	select {
		case msg1 := <-channel_1:
			fmt.Println("Received:", msg1)
		case msg2 := <-channel_2:
			fmt.Println("Received:", msg2)
	}

	// Watch out for closing channel when we have asynchronous execution ongoing...
	// Try to remove defer from channel_2
	defer close(channel_1)
	defer close(channel_2)

	// Select with default case (non-blocking)
	channel_3 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		channel_3 <- 10
	}()

	select {
		case msg := <-channel_3:
			fmt.Println("Received from channel_3:", msg)
		default:
			fmt.Println("No message received from channel_3 yet")
	}

	// Using channel for synchronization
	done := make(chan bool)

	go func() {
		fmt.Println("Working...")
		time.Sleep(1 * time.Second)
		fmt.Println("Done!")
		done <- true
	}()

	fmt.Println("Waiting...")
	<-done
	fmt.Println("Proceeding...")
}
