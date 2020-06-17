package main

import (
	"fmt"
)

const a int16 = 27

// iota is related to constant block
const (
	_ = iota // ignore first value by assigning to blank identifier
	x
	y
	z
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	x1 = iota
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PT
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)

	fmt.Printf("%v\n", x1)

	filesize := 4000000000.

	fmt.Printf("%.2fGB\n", filesize/GB)

	const myConst int = 42
	const a int = 14
	const b = 15

	var number int16 = 25

	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", number+b, number+b)
}
