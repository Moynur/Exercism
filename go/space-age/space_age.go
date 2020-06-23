package space

// Planet hello
type Planet string

var (
	secondToYears float64 = 31557600
	output                = ""
)

// Age hello
func Age(input float64, planet Planet) float64 {
	orbitalPeriodInEarthYears := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	result := input / (secondToYears * orbitalPeriodInEarthYears[planet])
	return result
}
