package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "مهدی خانزادی"

	l := len(s)
	c := utf8.RuneCountInString(s)

	fmt.Printf("number of bytes: %d, rune count: %d", l, c)
}
