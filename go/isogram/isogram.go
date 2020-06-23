package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram hello
// IsIsogram returns isogram bool
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i, letters := range word {
		if unicode.IsLetter(letters) && strings.ContainsRune(word[i+1:], letters) {
			return false
		}
	}
	return true
}
