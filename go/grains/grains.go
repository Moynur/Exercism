package grains

import (
	"errors"
)

// Square will square 2 up to the power of 63
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("can't square 0 or less or more than 63")
	}
	return 1 << (input - 1), nil
}

// Total returns 2^64
func Total() uint64 {
	return 1<<64 - 1
}
