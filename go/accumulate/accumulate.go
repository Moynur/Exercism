package accumulate

// Accumulate hello
func Accumulate(input []string, converter func(string) string) []string {
	var output = make([]string, len(input))
	for i, entry := range input {
		output[i] = converter(entry)
	}
	return output
}
