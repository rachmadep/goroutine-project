package goroutine_project

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			for range 100 {
				x = x + 1
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