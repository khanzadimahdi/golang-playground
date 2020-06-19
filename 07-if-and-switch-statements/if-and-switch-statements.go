package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("the test is true")
	}

	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}

	if pop, ok := statePopulation["Florida"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 30

	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too Hight")
	} else {
		fmt.Println("Equal :D")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// tagless syntax
	j := 10

	switch {
	case j <= 10:
		fmt.Println("less than ten")

		fallthrough
	case j <= 20:
		fmt.Println("less than 20")
	case j >= 20:
		fmt.Println("greater than 20")
	}

	var i interface{} = 1

	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
		break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is an string")
	case [3]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another type")
	}
}
