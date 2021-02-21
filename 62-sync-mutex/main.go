package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock() // a deferred Unlock will run even if the critical section panics

	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock() // a deferred Unlock will run even if the critical section panics

	return balance
}

func main() {
	Deposit(1000)
	fmt.Println(Balance())
}
