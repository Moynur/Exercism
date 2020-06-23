package diffsquares

// SquareOfSum returns the square of the sum of all natural numbers up to input
func SquareOfSum(input int) int {
	sum := input * (input + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of all squares before and including a given natural number
func SumOfSquares(input int) int {
	squares := input * (input + 1) * (2*input + 1) / 6
	return squares
}

// Difference comment
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
