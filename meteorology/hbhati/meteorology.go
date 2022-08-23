package meteorology

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tunit TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tunit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {
	return fmt.Sprintf("%v %v",
		temp.degree,
		temp.unit,
	)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (sunit SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sunit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (spd Speed) String() string {
	return fmt.Sprintf("%v %v",
		spd.magnitude,
		spd.unit,
	)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (mdata MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		mdata.location,
		mdata.temperature.String(),
		mdata.windDirection,
		mdata.windSpeed.String(),
		mdata.humidity,
	)
}
