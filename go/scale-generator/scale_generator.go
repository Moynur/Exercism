package scale

import "strings"

var (
	set1 = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	set2 = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

// Scale comment
func Scale(tonic, interval string) []string {
	scale := []string{}
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scale = set1
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = set2
	}
	tonic = strings.Title(tonic)
	for i, elem := range scale {
		if elem == tonic {
			scale = append(scale[i:], scale[:i]...)
			break
		}
	}

	if interval == "" {
		return scale
	}
	partialScale := []string{}
	stepsizeMap := map[string]int{"m": 1, "M": 2, "A": 3}
	i := 0
	for _, IntervalLength := range strings.Split(interval, "") {
		step := stepsizeMap[IntervalLength]
		partialScale = append(partialScale, scale[i%len(scale)])
		i += step
	}
	return partialScale
}
