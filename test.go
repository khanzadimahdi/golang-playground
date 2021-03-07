package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	y := reflect.ValueOf(&x)

	if y.CanAddr() {
		z := y.Elem()
		fmt.Println(z)
	}

	fmt.Println(x)
}
