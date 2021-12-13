package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "John Doe")

	got := buffer.String()
	want := "Hello, John Doe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
