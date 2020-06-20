package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(1)

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	waitGroup.Add(2)
	go sayHello()
	// time.Sleep(100 * time.Microsecond)

	var msg = "hello"

	go func() {
		fmt.Println(msg)

		waitGroup.Done()
	}()
	msg = "Goodbye"

	waitGroup.Wait()

	// time.Sleep(100 * time.Microsecond)
}

func sayHello() {
	fmt.Println("hello")

	waitGroup.Done()
}
