package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

func (tc TemperatureUnit) String() string {
	if tc == 0 {
		return "°C"
	} else {
		return "°F"
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (rc Temperature) String() string {
	return fmt.Sprintf("%d %v", rc.degree, rc.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (sc SpeedUnit) String() string {
	if sc == 0 {
		return "km/h"
	} else {
		return "mph"
	}

}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%d %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (mt MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %d%% Humidity", mt.location, mt.temperature, mt.windDirection, mt.windSpeed, mt.humidity)
}
