package main

import (
	"fmt"
	"strconv"
)

var z int = 42

var (
	firstName string = "Mahdi"
	lastName  string = "khanzadi"
	age       int    = 20
)

var i int = 100

func main() {
	fmt.Println(firstName)

	// format 1:
	var i int
	i = 42

	// format 2:
	var j int = 42

	// format 3:
	k := 42

	// format:
	var l = float32(j)

	fmt.Println(i, j, k)
	fmt.Println(string(k))
	fmt.Println(strconv.Itoa(k))

	fmt.Printf("%v, %T \n", k, k)

	fmt.Printf("%v, %T", l, l)
}
