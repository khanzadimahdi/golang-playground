package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int

	go func() {
		x = 1
		fmt.Printf("y:%d	", y)
	}()

	go func() {
		y = 1
		fmt.Printf("x:%d	", x)
	}()

	time.Sleep(1 * time.Second)
}

/**
	Notice:
	in absence of explicit synchronization, the compiler and CPU are
	free to reorder accesses to memory in any number of ways, so long as the
	behavior of each goroutine is sequentially consistent.
**/
