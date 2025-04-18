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

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()

	account.Balance += amount
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	defer account.RWMutex.RUnlock()

	return account.Balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(4 * time.Second)

	fmt.Println("Final balance:", account.GetBalance())
}