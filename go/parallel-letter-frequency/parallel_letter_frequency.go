package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var chanBuffer = 10

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculate frequency
func ConcurrentFrequency(list []string) FreqMap {
	c := make(chan FreqMap, chanBuffer)
	m := FreqMap{}
	for _, v0 := range list {
		go func(v0 string) {
			c <- Frequency(v0)
		}(v0)
	}

	for range list {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
