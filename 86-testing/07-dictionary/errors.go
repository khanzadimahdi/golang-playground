package dictionary

var ErrNotFound = DictionaryErr("could not find the word you were looking for")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
