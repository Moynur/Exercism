// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"reflect"
	"regexp"
	"strings"
)

var (
	response     = ""
	stringType   = reflect.TypeOf((""))
	isLetters, _ = regexp.Compile("[a-zA-Z]")
)

// Hey hello
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	switch {
	case len(remark) == 0:
		response = "Fine. Be that way!"
	case strings.ToUpper(remark) == remark && strings.HasSuffix(remark, "?") && isLetters.MatchString(remark):
		response = "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		response = "Sure."
	case strings.ToUpper(remark) == remark && !strings.HasSuffix(remark, "?") && isLetters.MatchString(remark):
		response = "Whoa, chill out!"
	default:
		response = "Whatever."
	}
	return response
}
