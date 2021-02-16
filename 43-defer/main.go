package main

import "fmt"

func main() {
	double(3)
	fmt.Println(plusTwo(2))
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()

	return x + x
}

func plusTwo(x int) (result int) {
	defer func() {
		result += 2
	}()

	return x
}
