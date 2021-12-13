package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "مهدی خانزادی"

	// strategy 1
	for i := 0; i < len(text); {
		r, s := utf8.DecodeRuneInString(text[i:])

		fmt.Printf("%c", r)

		i += s
	}

	fmt.Println("\n" + strings.Repeat("#", 30))

	// strategy 2
	for _, r := range text {
		fmt.Printf("%c", r)
	}
}
