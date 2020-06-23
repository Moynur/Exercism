package pythagorean

// Triplet blah
type Triplet [3]int

// Range comment
func Range(min, max int) []Triplet {
	var triples []Triplet
	for n := 1; n <= max; n++ {
		for m := n + 1; m <= max; m++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			if a < min || b < min || c < min || a > max || b > max || c > max {
			} else {
				if a > b {
					a, b = b, a
					triples = append(triples, Triplet{a, b, c})
				} else {
					triples = append(triples, Triplet{a, b, c})
				}

			}
		}

	}
	return triples
}

// Sum returns a b and c in a triplet for which a + b + c = total
func Sum(total int) []Triplet {
	var triples []Triplet
	for a := 1; a <= total; a++ {
		for b := a + 1; b <= total; b++ {
			c := total - b - a
			if a*a+b*b == c*c {
				triples = append(triples, Triplet{a, b, c})
			}
		}
	}
	return triples
}
