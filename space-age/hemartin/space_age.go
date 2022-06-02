package space

type Planet string

const orbitalPeriodEarh = 31557600

var secondsPerYear = map[Planet]float64{
	Planet("Mercury"): orbitalPeriodEarh * 0.2408467,
	Planet("Earth"):   orbitalPeriodEarh,
	Planet("Venus"):   orbitalPeriodEarh * 0.61519726,
	Planet("Mars"):    orbitalPeriodEarh * 1.8808158,
	Planet("Jupiter"): orbitalPeriodEarh * 11.862615,
	Planet("Saturn"):  orbitalPeriodEarh * 29.447498,
	Planet("Uranus"):  orbitalPeriodEarh * 84.016846,
	Planet("Neptune"): orbitalPeriodEarh * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsPerYear[planet]
}
