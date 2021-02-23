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
