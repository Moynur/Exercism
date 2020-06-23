package hamming

import "errors"

// Distance comment
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("inputs not same length")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}