package main

import "fmt"

func init() {
	fmt.Println("app/entry.go ==> init()")
}

var myVersion = fetchVersion()

func init() {
	defer fmt.Println("defer entry.go ==> init()")
	fmt.Println("entry.go ==> init()")
}

func main() {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	fmt.Println("version ===> ", myVersion)
}
