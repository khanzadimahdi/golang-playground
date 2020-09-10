package version

import "fmt"

func init() {
	fmt.Println("version/get-version.go ==> init()")
}

func getVersion() string {
	fmt.Println("version/get-version.go ==> getVersion()")
	return "1.0.0"
}
