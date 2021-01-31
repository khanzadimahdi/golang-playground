package main

import "fmt"

func main() {
	a := [...]int{1, 1, 2, 3, 5, 8, 13, 23, 36, 59}

	reverse(a[:])

	fmt.Println(a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// [i --> ... <- j]
