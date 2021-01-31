package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	wordsCount := make(map[string]int)

	for scanner.Scan() {
		wordsCount[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	for w, i := range wordsCount {
		fmt.Printf("%-10q\t%d\n", w, i)
	}
}
