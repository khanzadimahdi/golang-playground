package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a

	fmt.Println(a, *b)

	a = 27

	fmt.Println(a, *b)

	*b = 14

	fmt.Println(a, *b)

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]

	fmt.Printf("%v %p %p\n", c, d, e)

	var ms *myStruct

	fmt.Println(ms)

	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var ams *myStruct
	ams = new(myStruct)

	(*ams).foo = 55
	fmt.Println((*ams).foo)
}

type myStruct struct {
	foo int
}
