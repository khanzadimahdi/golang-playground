package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()

	counter()
}

func printStack() {
	var buf [4096]byte

	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func counter() {
	for i := 1; i < 10; i++ {
		i := i

		defer fmt.Println(i)
	}

	panic("end")
}
