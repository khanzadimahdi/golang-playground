package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}

	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (writer ConsoleWriter) Write(data []byte) (int, error) {
	n, error := fmt.Println(string(data))

	return n, error
}
