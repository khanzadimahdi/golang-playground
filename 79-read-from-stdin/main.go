package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%T\n", scanner)

	fmt.Print("write yourname:")

	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	text := scanner.Text()
	bytes := scanner.Bytes()

	log.Printf("text:%s\n", text)
	log.Printf("bytes:%b\n", bytes)
}
