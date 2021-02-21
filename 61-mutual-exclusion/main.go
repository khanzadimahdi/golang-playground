package main

import "fmt"

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema //release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token

	return b
}

func main() {
	Deposit(100)
	fmt.Println(Balance())
}
