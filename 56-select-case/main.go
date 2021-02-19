package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("%-3.d", x)
		case ch <- i:
		}
	}
}
