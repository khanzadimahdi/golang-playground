package text

import (
	"io/ioutil"
	"os"
	"unicode/utf8"
)

func CountChars(path string) (int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return utf8.RuneCount(data), nil
}

func FileLen(f string, bufsize int64) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	buf := make([]byte, bufsize)
	for {
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
}
