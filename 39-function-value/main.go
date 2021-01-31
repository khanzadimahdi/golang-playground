package main

import "fmt"

func sum(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func squares() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	result := sum(5, 10, 15, 20)

	fmt.Println(result)

	var f func(...int) int
	if f == nil { // zero value of a func type is nil
		fmt.Println("f is nil")
	}

	f = sum

	fmt.Println(f(5, 10, 15, 20))

	s := squares()

	fmt.Println(s()) // 1
	fmt.Println(s()) // 4
	fmt.Println(s()) // 9
}
