package main

import "fmt"

func main() {
	var any interface{}

	any = "mahdi"
	any = 123
	any = new([]byte)

	fmt.Println(any)
}
