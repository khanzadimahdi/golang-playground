package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	defer fmt.Println("start-1")
	defer fmt.Println("middle-1")
	defer fmt.Println("end-1")

	a := "start-02"

	defer fmt.Println(a)

	a = "end-02"

	fmt.Println("start-03")
	defer fmt.Println("this was differred")

	panicker()

	fmt.Println("end")
}

func panicker() {
	defer func() {
		if error := recover(); error != nil {
			log.Println("Error:", error)
		}
	}()

	panic("something bad happened")
}
