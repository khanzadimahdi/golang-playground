package version

import "fmt"

func init() {
	fmt.Println("version/entry.go ==> init()")
}

var Version = getLocalVersion()

func getLocalVersion() string {
	fmt.Println("version/entry.go ==> getLocalVersion()")
	return getVersion()
}
