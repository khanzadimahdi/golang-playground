package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press Enter to abort.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // wait to read a single byte
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return // cause ticker goroutine leak
		}
	}

	launch()
}

func launch() {
	fmt.Println("launched...")
}
