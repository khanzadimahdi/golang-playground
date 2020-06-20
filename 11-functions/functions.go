package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world")

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()

	fmt.Println("The new name is:", g.name)

	// for i := 0; i < 5; i++ {
	// 	sayMessage("hello Go!", i)
	// }

	// greeting := "hello"
	// name := "stacey"

	// sayGreeting(&greeting, &name)

	// fmt.Println(name)

	// sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// fmt.Println(sumResult(1, 2, 3, 4, 5, 6, 7, 8, 9))

	d, error := divide(5.0, 3.0)

	if error != nil {
		fmt.Println(error)

		return
	}

	fmt.Println(d)

	func() {
		fmt.Println("Immediately invoked function expression")
	}()

	f := func() {
		fmt.Println("Hello Go")
	}

	f()

	var j func() = func() {
		fmt.Println("Hello Go")
	}

	j()
}

func sum(values ...int) {
	fmt.Println(values)

	result := 0

	for _, v := range values {
		result += v
	}

	fmt.Println("The sum is", result)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg, idx)
}

func sayGreeting(greeting, name *string) {
	fmt.Println("greeting", *name)
	*name = "Ted"
	fmt.Println(*name)
}

func anotherSum(values ...int) *int {
	fmt.Println(values)
	result := 0

	for _, v := range values {
		result += v
	}

	return &result
}

func sumResult(values ...int) (result int) {
	fmt.Println(values)

	for _, v := range values {
		result += v
	}

	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
