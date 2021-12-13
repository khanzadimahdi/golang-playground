package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// type of scanning
	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("line of file:", scanner.Text())
	}
}
