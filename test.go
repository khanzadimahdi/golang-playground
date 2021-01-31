package main

import "fmt"

func main() {
	factorial := func(number int) int {
		if number == 1 {
			return 1
		}

		return number * factorial(number-1) // compile error undefined: factorial
	}

	fmt.Println(factorial(4))
}
