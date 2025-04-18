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

type UserBalance struct {
	Mutex sync.Mutex
	Name  string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Locking user1:", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)
	
	user2.Lock()
	fmt.Println("Locking user2:", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := &UserBalance{Name: "User1", Balance: 10000}
	user2 := &UserBalance{Name: "User2", Balance: 10000}

	go Transfer(user1, user2, 100)
	go Transfer(user2, user1, 100)

	time.Sleep(5 * time.Second)

	fmt.Println("Final balance of user1:", user1.Balance)
	fmt.Println("Final balance of user2:", user2.Balance)
}