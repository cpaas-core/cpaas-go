package space

var NumSecondsInEarthYear float64 = 31557600.0

var OrbitalPeriod = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

type Planet string

func (p *Planet) OrbitalPeriod() float64 {
	return OrbitalPeriod[*p]
}

func (p *Planet) SecondsInAYear() float64 {
	return p.OrbitalPeriod() * NumSecondsInEarthYear
}

func (p *Planet) HowOldAmI(seconds float64) float64 {
	return seconds / p.SecondsInAYear()
}

func Age(seconds float64, planet Planet) float64 {
	return planet.HowOldAmI(seconds)
}
