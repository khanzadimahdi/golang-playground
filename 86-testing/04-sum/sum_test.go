package sum

import (
	"reflect"
	"testing"
)

func sumAll(numbersToSum ...[]int) []int {
	l := len(numbersToSum)
	sums := make([]int, l)

	for i, n := range numbersToSum {
		sums[i] = Sum(n)
	}

	return sums
}

func TestSum(t *testing.T) {
	got := sumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	want := []int{15, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
