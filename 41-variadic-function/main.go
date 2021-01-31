package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

func sum(vals ...int) (result int) {
	for _, val := range vals {
		result += val
	}

	return
}
