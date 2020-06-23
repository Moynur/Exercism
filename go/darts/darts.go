package darts

import "math"

// Score hello
func Score(x float64, y float64) int {
	points := 0
	radius := math.Sqrt(x*x + y*y)
	if radius <= 1 {
		points = 10
	} else if radius <= 5 {
		points = 5
	} else if radius <= 10 {
		points = 1
	} else {
		points = 0
	}
	result := points
	return result
}
