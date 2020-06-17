package main

import "fmt"

func main() {
	/**
	working with booleans
	*/

	var isActive bool = false

	m := 1 == 1
	n := 1 == 10

	var test bool

	fmt.Printf("%v, %T\n", isActive, isActive)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", test, test)

	/**
	working with integers
	*/

	var number int // at least 32 bit number
	var number8bit int8
	var number16bit int16
	var number32bit int32
	var number64bit int64

	fmt.Printf("%v, %T\n", number, number)
	fmt.Printf("%v, %T\n", number8bit, number8bit)
	fmt.Printf("%v, %T\n", number16bit, number16bit)
	fmt.Printf("%v, %T\n", number32bit, number32bit)
	fmt.Printf("%v, %T\n", number64bit, number64bit)

	var numberUnsigned8bit uint8
	var numberUnsigned16bit uint16
	var numberUnsigned32bit uint32
	// we don't have 64 bit unsigned integer

	fmt.Printf("%v, %T\n", numberUnsigned8bit, numberUnsigned8bit)
	fmt.Printf("%v, %T\n", numberUnsigned16bit, numberUnsigned16bit)
	fmt.Printf("%v, %T\n", numberUnsigned32bit, numberUnsigned32bit)

	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // and 0010
	fmt.Println(a | b)  // or 1011
	fmt.Println(a ^ b)  // xor 1001
	fmt.Println(a &^ b) // and-not // 0100

	x := 8 // 2^3

	fmt.Println(x << 3) // shift 3 places to left 2^3 * 2^3 = 2^6 = 64
	fmt.Println(x >> 3) // shift 3 places to right 2^3 / 2^3 = 2^0 = 1

	/**
	working with float numbers
	*/

	var float32bit float32 = 3.14
	var float64bit float64 = 13.7e72

	fmt.Printf("%v,%T \n", float32bit, float32bit)
	fmt.Printf("%v,%T \n", float64bit, float64bit)

	floatNumber1 := 10.2
	floatNumber2 := 3.7

	fmt.Println(floatNumber1 + floatNumber2)
	fmt.Println(floatNumber1 - floatNumber2)
	fmt.Println(floatNumber1 * floatNumber2)
	fmt.Println(floatNumber1 / floatNumber2)

	/**
	working with complex numbers
	*/

	var complex64bitNumber complex64 = 1.2 + 2.1
	var complex128bitNumber complex128 = 1 + 2i

	fmt.Printf("%v,%T \n", complex64bitNumber, complex64bitNumber)
	fmt.Printf("%v,%T \n", complex128bitNumber, complex128bitNumber)

	fmt.Printf(
		"real part %v, imag part %v",
		real(complex128bitNumber),
		imag(complex128bitNumber),
	)

	var xyz complex128 = complex(5, 12)

	fmt.Printf("%v, %T\n", xyz, xyz)

	/**
	working with strings
	*/
	simpleString := "this is a string"
	convertiontoBytes := []byte(simpleString)

	fmt.Printf("%v, %T\n", simpleString, simpleString)
	fmt.Printf("%v, %T\n", string(simpleString[2]), simpleString[2])
	fmt.Printf("%v, %T\n", convertiontoBytes, convertiontoBytes)

	fmt.Println(simpleString + "cancat with another string")

	// rune : alias for type int32
	r := 'a'
	var s rune = 'a'

	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", s, s)
}
