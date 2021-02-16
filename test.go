package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("in goroutine")

		done <- 1
	}()

	fmt.Println(<-done)
}
