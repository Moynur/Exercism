package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var sentence = ""
	if name != "" {
		sentence = "One for " + name + ", one for me."
	} else {
		sentence = "One for you, one for me."
	}
	return sentence
}
