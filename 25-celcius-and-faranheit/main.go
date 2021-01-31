package main

import (
	"fmt"

	"./tempConv"
)

func main() {
	temprature := tempConv.Celsius(24)

	fmt.Println(temprature.ToFaranheit())
}
