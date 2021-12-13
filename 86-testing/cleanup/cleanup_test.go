package cleanup

import (
	"os"
	"testing"
)

// createFile is a helper function called from multiple tests

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, dont worry about cleanup
}
