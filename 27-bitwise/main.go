package main

import "fmt"

func main() {
	var x int8 = 1 << 1
	var y int8 = 1 << 5

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b", x&^y)

	fmt.Printf("%d", 0xdeadbeaf)
}
