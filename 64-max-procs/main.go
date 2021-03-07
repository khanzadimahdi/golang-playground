package main

import (
	"fmt"
)

func main() {
	// runtime.GOMAXPROCS(2)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

// GOMAXPROCS=2 go run main.go
