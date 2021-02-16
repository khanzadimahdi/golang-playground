package main

import "fmt"

type BytesCounter int

func (c *BytesCounter) Write(p []byte) (n int, err error) {
	*c += BytesCounter(len(p)) // Convert int to ByteCounter
	return len(p), nil
}

func main() {
	lorem := "lorem ipsum"

	var counter BytesCounter

	fmt.Fprintf(&counter, lorem)

	fmt.Print(counter)
}
