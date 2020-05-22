package service

// WordList is a simple type that holds lists of words
type WordList []string

func (w WordList) contains (e string) bool {
	for _, item := range(w) {
		if item == e {
			return true
		}
	}
	return false
}