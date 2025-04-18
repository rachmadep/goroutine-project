package goroutine_project

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done() // Decrement the counter when the goroutine completes

	group.Add(1) // Increment the counter before starting the goroutine

	fmt.Println("Asynchronous task completed")
	time.Sleep(1 * time.Second) // Simulate some work
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsyncronous(group)
	}

	group.Wait() // Wait for all goroutines to finish
	fmt.Println("All asynchronous tasks completed")
}