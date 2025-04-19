package goroutine_project

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1) // Use atomic operation to increment x
			}
			group.Done()
		}()
	}

	group.Wait() // Wait for all goroutines to finish
	// Check the value of x
	fmt.Println("Final value of x:", x)
}