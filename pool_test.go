package goroutine_project

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// New: func() interface{} {
		// 	return "New Itema"
		// },
	}
	group := sync.WaitGroup{}

	pool.Put("Hello, Pool!")
	pool.Put("Hello, Pool 2!")
	pool.Put("Hello, Pool 3!")

	for i := 0; i < 20; i++ {
		go func() {
			group.Add(1)
			item := pool.Get()
			fmt.Println(item)
			time.Sleep(1 * time.Second)

			pool.Put(item)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("All tasks completed")
}