package raindrops

import "strconv"

// Convert hello
func Convert(input int) string {
	result := ""
	if input%3 == 0 {
		result = result + "Pling"
	}
	if input%5 == 0 {
		result = result + "Plang"
	}
	if input%7 == 0 {
		result = result + "Plong"
	}
	if input%7 != 0 && input%5 != 0 && input%3 != 0 {
		result = strconv.Itoa(input)
	}
	return result
}
