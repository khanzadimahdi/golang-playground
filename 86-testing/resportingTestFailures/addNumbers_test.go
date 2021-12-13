package numbers

import "testing"

func Test_addNumbers(t *testing.T) {
	x := 1
	y := 2
	z := 3

	if r := addNumbers(x, y); r != z {
		t.Errorf("incorrect result: expecting %d got %d\n", z, r)
	}
}
