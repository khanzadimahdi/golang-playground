package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("want: '%d', got: '%d'", want, got)
	}
}
