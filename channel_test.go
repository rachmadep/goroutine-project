package goroutine_project

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		// time.Sleep(2 * time.Second) // Simulate some work
		channel <- "Hello, Channel!"
		fmt.Println("Message sent to channel")
	}()

	// Receive the message from the channel
	message := <- channel
	fmt.Println("Received message:", message)

	time.Sleep(2 * time.Second) // Wait for the goroutine to finish
}