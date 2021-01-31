package main

import "fmt"

func main() {
	var ascii = 'a'
	var newline = '\n'

	fmt.Printf("%c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}
