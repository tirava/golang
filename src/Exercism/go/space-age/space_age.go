// Implements calculating orbital period in Earth years
package space

type Planet string

var ages = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates period from seconds
func Age(seconds float64, planet Planet) float64 {
	return seconds / 3600 / 24 / 365.25 / ages[planet]
}
