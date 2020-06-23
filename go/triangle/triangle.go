package triangle

// Kind hello
type Kind string

var (
	k Kind
)

const (
	// NaT hello
	NaT Kind = "Not a triagnle"
	// Equ hello
	Equ Kind = "Equilateral triangle"
	// Iso hello
	Iso Kind = "iso"
	// Sca hello
	Sca Kind = "Scalence"
)

func checkIsTriagnel(a, b, c float64) Kind {
	switch {
	case a+b < c:
		k = NaT
	case a+c < b:
		k = NaT
	case b+c < a:
		k = NaT
	case a == 0:
		k = NaT
	case b == 0:
		k = NaT
	case c == 0:
		k = NaT
	}
	return k
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if checkIsTriagnel(a, b, c) == NaT {
		return NaT
	}
	switch {
	case a == b && b == c:
		k = Equ
	case a == b && b != c:
		k = Iso
	case a != b && b == c:
		k = Iso
	case a == c && c != b:
		k = Iso
	case a != b && b != c && a != c:
		k = Sca
	}
	return k
}
