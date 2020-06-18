package main

import (
	"fmt"
	"reflect"
)

type doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required:"value" max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})

	field, _ := t.FieldByName("Name")

	fmt.Println(field.Tag)

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	fmt.Println(statePopulations)

	statePopulations1 := make(map[string]int)

	statePopulations1 = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	fmt.Println(statePopulations1)

	fmt.Println(statePopulations1["California"])

	statePopulations1["Georgia"] = 10310371
	fmt.Println(statePopulations1["Georgia"])

	delete(statePopulations1, "Georgia")
	fmt.Println(statePopulations1["Georgia"])

	_, ok := statePopulations1["Georgia"]
	fmt.Println(ok)

	fmt.Println(len(statePopulations1))

	aDoctor := doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)

	fmt.Println(aDoctor.actorName)

	anotherDoctor := doctor{
		3,
		"Jon Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(anotherDoctor)
}
