package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File

	fmt.Println("%T/n", newFile)

	var err error

	// creates a file if it not exists
	// if it exists, it will truncated.
	newFile, err = os.Create("a.txt")

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)

		log.Fatal(err)
	}

	fmt.Printf("%#v\n", newFile)

	// truncate given file to 0 byte
	err = os.Truncate("a.txt", 0)

	if err != nil {
		log.Fatal(err)
	}

	// when we are done with the file, we must close it
	newFile.Close()

	// open file in readonly mode
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()

	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo

	fileInfo, err = os.Stat("a.txt")

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file is not exist")
		}
		fmt.Println(err)
	}

	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file last modified:", fileInfo.ModTime())
	fmt.Println("file is directory:", fileInfo.IsDir())
	fmt.Println("file permission:", fileInfo.Mode())

	oldPath := "a.txt"
	newPath := "renamedFile.txt"

	// rename or move file
	err = os.Rename(oldPath, newPath)

	if err != nil {
		log.Fatal(err)
	}

	// remove a file
	fileNames := []string{
		"renamedFile.txt",
		"test.txt",
	}

	for _, name := range fileNames {
		log.Println("try to remove file:", name)

		err = os.Remove(name)

		if err != nil {
			log.Fatal(err)
		}
	}
}
