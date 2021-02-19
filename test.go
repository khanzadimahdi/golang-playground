package main

import (
	"fmt"
	"time"
)

func main() {
loop:
	for {
		x := 9
		time.Sleep(1 * time.Second)
		x++
		fmt.Println(x)
		if x >= 10 {
			break loop
		}
	}
	fmt.Println("after loop")
}
