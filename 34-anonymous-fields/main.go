package main

import "fmt"

func main() {
	namedFields()
	anonymousFields()
}

func namedFields() {
	type Point struct {
		X int
		Y int
	}

	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}

	var w Wheel

	w.Circle.Center.X = 0
	w.Circle.Center.Y = 0
	w.Circle.Radius = 20
	w.Spokes = 5

	fmt.Printf("%#v", w)
}

func anonymousFields() {
	type Point struct {
		X int
		Y int
	}

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel

	w.X = 0
	w.Y = 0
	w.Radius = 20
	w.Spokes = 5

	fmt.Printf("%#v", w)
}

// notice 1:
//-------------------------------------
// A. because "anonymous" fields do have implicit names, you can't have two anonymous
// fields of the same type since their names whould conflict.
//-------------------------------------
// B. and because the name of the field is implicitly determined by its type,
// so too is visibility of field.
//-------------------------------------

// notice 2:
// the # adverb caused Printf's %v verb to display
// values in a form similar to Go syntax
// for struct values, this form includes the name of each field.
