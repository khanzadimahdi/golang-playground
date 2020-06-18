package main

import (
	"fmt"
)

func main() {
	grade1 := 20
	grade2 := 19
	grade3 := 17.5

	fmt.Printf("grades are: %v, %v, %v \n", grade1, grade2, grade3)

	grades := [...]int{97, 85, 93}

	fmt.Printf("Grades: %v \n", grades)

	var students [3]string

	students[0] = "Alex Zarkov"
	students[1] = "John Doe"
	students[2] = "Mahdi Khanzadi"

	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Student #1: %v \n", students[1])
	fmt.Printf("Number of Students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}

	fmt.Printf("identity martix: %v \n", identityMatrix)

	a := [...]int{1, 2, 3}

	b := a

	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)

	// pointing to the same data

	c := [...]int{1, 2, 3}
	d := &c

	d[1] = 5

	fmt.Println(c)
	fmt.Println(d)

	/**
	working with slices
	*/
	i := []int{1, 2, 3}

	fmt.Printf("slice value: %v", i)
	fmt.Printf("slice length: %v", len(i))
	fmt.Printf("slice capacity: %v", cap(i))

	// slices are assign by address
	j := i

	j[1] = 20

	fmt.Println(i)
	fmt.Println(j)

	k := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// [inclusive, exclusive]
	// or we can say ==> [from, up to]

	l := k[:]   // slice of all elements
	m := k[3:]  // slice from 4th element to end
	n := k[:6]  // slice first 6 elements
	o := k[3:6] // slice the 4th, 5th and 6th elements

	k[5] = 42

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

	p := make([]int, 3, 100)

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(p))

	q := []int{}

	fmt.Println(q)
	fmt.Printf("Length: %v\n", len(q))
	fmt.Printf("Capacity: %v\n", cap(q))

	q = append(q, 1)

	fmt.Println(q)
	fmt.Printf("Length: %v\n", len(q))
	fmt.Printf("Capacity: %v\n", cap(q))

	q = append(q, 2, 3, 4, 5)

	fmt.Println(q)
	fmt.Printf("Length: %v\n", len(q))
	fmt.Printf("Capacity: %v\n", cap(q))

	q = append(q, []int{6, 7, 8, 9, 10}...)

	fmt.Println(q)
	fmt.Printf("Length: %v\n", len(q))
	fmt.Printf("Capacity: %v\n", cap(q))

	r := [...]int{1, 2, 3, 4, 5}

	fmt.Println(r)

	s := append(r[:2], r[3:]...)

	s[0] = 10

	fmt.Println(s)
	fmt.Println(r)

	var t [3]int = [3]int{1, 2, 3}

	fmt.Println(t)
}
