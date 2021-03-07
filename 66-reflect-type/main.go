package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)

	fmt.Println(t.String()) // a reflect.Type
	fmt.Println(t)

	v := reflect.ValueOf(3)
	fmt.Println(v.Type())

	fmt.Printf("%T\n", 3) // uses reflect.TypeOf internally
}
