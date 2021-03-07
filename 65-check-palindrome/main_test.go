package main

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Able was I ere I saw Elba", true},
		{"palindrome", false}, //non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

// run: go test -v
// run selective: go test -v -run="French|Canal"
//............................................................................
// 1. table-driven testing is very common in Go. (see tests struct)
//............................................................................
// 2. t.Errorf does not cause a panic or stop the execution of the test,
//    unlike assertion failures in many test frameworks exists for
//    other languages.
//............................................................................
// 3. when we want stop a test function we use t.Fatal or t.Fatalf
//	  this must be called from the same goroutine as the Test function,
//    not from another one created during the test.
//............................................................................
// 4. failure messages in a Test, has usually the "f(x) = y, want z" form
//    where f(x) explains the attempted operation and its input, y is the
//    actual result, and z the expected result.
//    messages must avoid boilerplate and redundant information.
//    when testing a boolean function, omit `want z` part since it adds
//    no information.
//............................................................................
// 5. its important the code under test not call log.Fatal or os.Exit, since
//    these will stop the proccess in its tracks.
//............................................................................
// 6. if something unexpected happens and a function panics, the test driver
//    will recover and the test will be considered as failure.
//............................................................................

// profiling:
// go test -cpuprofile=cpu.out
// go test -blockprofile=block.out
// go test -memprofile=mem.out
// analyz with : go tool pprof
