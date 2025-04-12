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

func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second) // Simulate some work
	channel <- "Hello, Channel!"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	// Receive the message from the channel
	message := <- channel
	fmt.Println("Received message:", message)

	time.Sleep(3 * time.Second) // Wait for the goroutine to finish
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second) // Simulate some work
	channel <- "Hello, Channel!"
}

func OnlyOut(channel <-chan string) {
	message := <- channel
	fmt.Println("Received message:", message)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	// Wait for the goroutines to finish
	time.Sleep(3 * time.Second)
}

// Buffered channels

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Hello, Channel 1!"
		channel <- "Hello, Channel 2!"
		// channel <- "Hello, Channel 3!"
	}()

	go func ()  {
		time.Sleep(1 * time.Second) // Simulate some work
		fmt.Println("Buffered channel size:", len(channel))
		fmt.Println("Buffered channel capacity:", cap(channel))

		fmt.Println("Received message:", <- channel)
		fmt.Println("Received message:", <- channel)
		// fmt.Println("Received message:", <- channel)
	}()

	time.Sleep(2 * time.Second) // Wait for the goroutines to finish

	fmt.Println("Done with buffered channel")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Message %d", i)
		}

		time.Sleep(1 * time.Second) // Simulate some work
		close(channel) // Close the channel when done
		fmt.Println("Channel closed")
	}()

	for data := range channel {
		fmt.Println("Received message:", data)
	}

	fmt.Println("Done with range channel")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <- channel1:
				fmt.Println("Received message from channel1:", data)
				counter++
			case data := <- channel2:
				fmt.Println("Received message from channel2:", data)
				counter++
		}

		if counter == 2 {
			break
		}
	}
	

	fmt.Println("Done with select channel")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <- channel1:
				fmt.Println("Received message from channel1:", data)
				counter++
			case data := <- channel2:
				fmt.Println("Received message from channel2:", data)
				counter++
			default:
				fmt.Println("No messages received, doing other work...")
		}

		if counter == 2 {
			break
		}
	}
	

	fmt.Println("Done with default select channel")
}