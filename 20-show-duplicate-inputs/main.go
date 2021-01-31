package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter something: ")

	for input.Scan() {
		if input.Text() == "exit" {
			break
		}

		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, count := range counts {
		if count > 1 {
			fmt.Println(line)
		}
	}
}
