package acronym

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	acr := ""
	s = strings.Replace(s, "_", "", -1)
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "   ", " ", -1)
	s = strings.ToUpper(s)
	arrayOfStrings := strings.Split(s, " ")
	fmt.Println(arrayOfStrings[:])
	for _, str := range arrayOfStrings {
		r, _ := utf8.DecodeRuneInString(str)
		acr = acr + string(r)
	}
	return acr
}
