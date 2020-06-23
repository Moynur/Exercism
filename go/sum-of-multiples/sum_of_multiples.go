package summultiples

// SumMultiples up to limit provided for given factors
func SumMultiples(limits int, divisors ...int) (sum int) {
	sum = 0
	if len(divisors) == 0 || limits < divisors[0] {
		return sum
	}
	for i := 1; i < limits; i++ {
		for _, v := range divisors {
			if v != 0 && i%v == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
