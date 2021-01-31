package main

import "fmt"

func main() {
	names := []string{"john", "alex", "sara", "mike", "joseph"}

	var actions []func()

	for _, name := range names {
		name := name
		actions = append(actions, func() {
			fmt.Println(name)
		})
	}

	for _, action := range actions {
		action()
	}
}
