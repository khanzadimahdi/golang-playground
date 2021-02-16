package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	strings := []string{"mahdi", "ali", "mohamamd", "ashkan", "saeed"}

	sort.Sort(StringSlice(strings))

	fmt.Println(strings)

	strings2 := []string{"mahdi", "ali", "mohamamd", "ashkan", "saeed"}

	sort.Strings(strings2)

	fmt.Println(strings2)
}
