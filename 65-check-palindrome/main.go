package main

import "unicode"

func main() {

}

// IsPalindrome reports whether s reads the same forward and backward.
func IsPalindrome(s string) bool {
	var letters = make([]rune, 0, len(s))

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
