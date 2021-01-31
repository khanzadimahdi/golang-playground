package main

import "fmt"

func main() {
	var x uint8 = 255
	var y int8 = -128

	fmt.Println(x, x+1, x*x)
	fmt.Println(y, y-1, y*y)

	const e = 2.71828
	fmt.Println(e + .2)
}
