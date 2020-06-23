package pangram

import (
	"strings"
	"unicode"
)

// IsPangram a
func IsPangram(sentence string) bool {
	sentecne = strings.ToLower(sentence)
	for i, letters := range sentence {
		if unicode.IsLetter(letters) && strings.ContainsRune(word[i+1:], letters) {
			return false
		}
	}
	return true

}
