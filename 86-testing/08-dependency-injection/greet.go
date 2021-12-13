package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, data string) {
	fmt.Fprintf(writer, "Hello, %s", data)
}
