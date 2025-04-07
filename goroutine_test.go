package goroutine_project

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutine(t *testing.T) {
	// Create a goroutine
	go RunHelloWorld()
	fmt.Println("Goroutine created")

	// Wait for the goroutine to finish
	// In a real-world scenario, you would use sync.WaitGroup or channels to wait for goroutines
	// Here we just sleep for a short time to allow the goroutine to finish
	time.Sleep(1 * time.Second)

	// Note: In a real test, you would not use sleep like this. This is just for demonstration.
	t.Log("Goroutine started")
}

func DisplayNumber(number int) {
	fmt.Println("Number:", number)
}

func TestManyGoroutine(t *testing.T) {
	// Create many goroutines
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	// Wait for a short time to allow goroutines to finish
	time.Sleep(5 * time.Second)
}