package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	channel := make(chan int)
	waitGroup.Add(2)
	go func() {
		i := <-channel
		fmt.Println(i)
		waitGroup.Done()
	}()

	go func() {
		i := 42
		channel <- i
		i = 27
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
