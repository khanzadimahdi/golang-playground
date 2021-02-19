package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

// how to build with different architechture:
// GOARCH=386 go build .
