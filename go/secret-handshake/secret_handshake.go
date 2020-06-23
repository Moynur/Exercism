package secret

// Handshake takes an int and compares the binary representation and returns the secret handshake
func Handshake(input uint) []string {
	var empty []string
	// for i := len(word) - 1; i >= 0; i-- {
	// 	fmt.Println(i)
	switch {
	case input&1 == 1:
		print("\n", input)
		empty = append(empty, "wink")
	case input&2 == 2:
		print("\n", input)
		empty = append(empty, "double blink")
	case input&4 == 4:
		print("\n", input)
		empty = append(empty, "close your eyes")
	case input&8 == 8:
		print("\n", input)
		empty = append(empty, "jump")

	case input&16 == 16:
		print("\n", input)
		empty = reverse(empty)
	}

	// case i == 2:
	// 	if word[i] == 1 {
	// 		empty = append(empty, "wink")
	// 	}

	return empty
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
