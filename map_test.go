package goroutine_project

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int) {
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			AddToMap(data, i)
		}(i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}