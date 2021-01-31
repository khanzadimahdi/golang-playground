package version

import "fmt"

func init() {
	defer fmt.Println("defer version/entry.go ==> init()")
	fmt.Println("version/entry.go ==> init()")
}

var Version = getLocalVersion()

func getLocalVersion() string {
	fmt.Println("version/entry.go ==> getLocalVersion()")
	return getVersion()
}
