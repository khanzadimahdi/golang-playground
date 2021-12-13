package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	var file *os.File
	var err error

	file, err = os.OpenFile("myfile.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	bytes := []byte{97, 98, 99}

	count, err := bufferedWriter.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("bytes written fo buffer (not file):", count)

	bytesAvailable := bufferedWriter.Available()

	log.Println("bytes available in buffer:", bytesAvailable)

	count, err = bufferedWriter.WriteString("This is my Golang playground")

	if err != nil {
		log.Fatal(err)
	}

	unflashedBufferSize := bufferedWriter.Buffered()

	log.Println("unflashed buffer size:", unflashedBufferSize)

	// remove buffered data
	// bufferedWriter.Reset(bufferedWriter)

	// store buffered data info file
	bufferedWriter.Flush()
}
