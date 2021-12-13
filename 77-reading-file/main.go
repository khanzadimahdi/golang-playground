package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var file *os.File
	var err error

	file, err = os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := make([]byte, 5)

	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("number of bytes read:", numberBytesRead)
	log.Printf("data: %s\n", byteSlice)

	// read all file---------------------------------------
	file, err = os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data: %s\n", data)

	// ----------------------------------------------------

	data, err = ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data: %s\n", data)
}
