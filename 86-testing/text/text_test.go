package text

import (
	"fmt"
	"testing"
)

func TestCountChars(t *testing.T) {
	total, err := CountChars("testdata/sample.txt")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if total != 118 {
		t.Error("Expected 118, got", total)
	}

	_, err = CountChars("testdata/no_file.txt")
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestFileLen(t *testing.T) {
	num, err := FileLen("testdata/sample.txt", 1)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if num != 126 {
		t.Errorf("Expected: %d, got %d", 126, num)
	}

	_, err = FileLen("testdata/no_file.txt", 1)
	if err == nil {
		t.Error("Expected and error")
	}
}

var blackhole int

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int64{1, 5, 35, 70, 110, 150, 200} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/sample.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
