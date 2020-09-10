package main

import (
	"./version"
	"fmt"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}

func fetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version.Version
}
