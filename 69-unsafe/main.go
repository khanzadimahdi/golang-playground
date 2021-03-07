package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var y float32

	y = 22.3

	z := unsafe.Sizeof(x)

	fmt.Printf("%T %[1]v", z)

	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

	*pb = 42

	fmt.Println(x.b) // "42"
}
