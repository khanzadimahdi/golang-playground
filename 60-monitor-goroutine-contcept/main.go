package main

import (
	"fmt"
	"time"

	"./bank"
)

func main() {
	go func() {
		bank.Deposit(100)
		fmt.Println(bank.Balance())
	}()

	go func() {
		bank.Deposit(150)
		fmt.Println(bank.Balance())
	}()

	go func() {
		bank.Deposit(200)
		fmt.Println(bank.Balance())
	}()

	go func() {
		bank.Deposit(250)
		fmt.Println(bank.Balance())
	}()

	time.Sleep(1 * time.Second)
}
