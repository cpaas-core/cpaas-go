package space

type Planet string

var earthOrbitalPeriodInSeconds float64 = 31557600.0

var orbitalPeriodInSeconds = map[Planet]float64{
	"Earth":   earthOrbitalPeriodInSeconds,
	"Mercury": 0.2408467 * earthOrbitalPeriodInSeconds,
	"Venus":   0.61519726 * earthOrbitalPeriodInSeconds,
	"Mars":    1.8808158 * earthOrbitalPeriodInSeconds,
	"Jupiter": 11.862615 * earthOrbitalPeriodInSeconds,
	"Saturn":  29.447498 * earthOrbitalPeriodInSeconds,
	"Uranus":  84.016846 * earthOrbitalPeriodInSeconds,
	"Neptune": 164.79132 * earthOrbitalPeriodInSeconds,
}

func Age(seconds float64, planet Planet) float64 {
	if value, ok := orbitalPeriodInSeconds[planet]; ok {
		return seconds / value
	}
	return 0.0
}
