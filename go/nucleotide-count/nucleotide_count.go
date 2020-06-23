package dna

import (
	"fmt"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts hello
func (d DNA) Counts() (Histogram, error) {
	d = DNA(strings.ToUpper(string(d)))
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		_, ok := h[nucleotide]
		if !ok {
			return nil, fmt.Errorf("Invalid nucleotide %v", nucleotide)
		}
		h[nucleotide]++
	}
	return h, nil
}
