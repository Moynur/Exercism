package leap

var leapYear bool

// IsLeapYear should have a comment documenting it.
func IsLeapYear(input int) bool {
	if input%400 == 0 {
		leapYear = true
	} else if input%100 == 0 {
		leapYear = false
	} else if input%4 == 0 {
		leapYear = true
	}
	return leapYear
}
