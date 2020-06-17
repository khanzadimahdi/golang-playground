package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHeadQuarters

	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	var roles byte = isAdmin | canSeeAfrica | canSeeAsia

	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("is HQ? %v", isHeadQuarters&roles == isHeadQuarters)
}
