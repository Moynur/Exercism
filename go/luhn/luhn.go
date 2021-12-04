package luhn

import (
	"fmt"
	"strings"
)

// Valid this function determines whether a string of numbers passes the luhn algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) < 2 {
		return false
	}
	var number int
	var total int
	double := len(input)%2 == 0
	for _, r := range input {
		number = int(r - '0')
		if number < 0 || number > 9 {
			return false
		}
		if double == true {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		total += number
		double = !double
	}
	return total%10 == 0
}
