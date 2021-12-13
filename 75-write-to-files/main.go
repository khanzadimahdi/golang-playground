package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var file *os.File
	var err error

	file, err = os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSLice := []byte("I learn Golang!")

	count, err := file.Write(byteSLice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("numbe of bytes written to file:", count)

	err = ioutil.WriteFile("test.txt", []byte("This is another way to write into file"), 0644)

	if err != nil {
		log.Fatal(err)
	}
}
