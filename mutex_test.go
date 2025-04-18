package goroutine_project

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	// Create a mutex
	mutex := sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// Lock the mutex before accessing x
				mutex.Lock()
				x = x + 1

				// Unlock the mutex after accessing x
				mutex.Unlock()
			}
		}()
	}

	// Wait for a short time to allow goroutines to finish
	// In a real-world scenario, you would use sync.WaitGroup or channels to wait for goroutines
	// Here we just sleep for a short time to allow the goroutines to finish
	time.Sleep(4 * time.Second)

	// Check the value of x
	fmt.Println("Final value of x:", x)
}