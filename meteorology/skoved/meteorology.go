package meteorology

import (
	"fmt"
	"reflect"
)

// The underlying value for types that implement Unit must be representable by int. Example: type NewUnit int
type Unit interface {
	UnitSlice() []string // the value of the Unit consts should correspond to their string value's position in the slice
	String() string      // ensure that all types that implement Unit are Stringers
}

// Convert a Unit to a string. This function expects that the underlying value of the unit can be represented as an int
func UnitString(unit Unit) string {
	ref := reflect.ValueOf(unit)
	// panic if the underlying value of unit cannot be represented as int
	if !ref.CanInt() {
		panic("The underlying value of types that implement Unit must be representable by int")
	}
	return unit.UnitSlice()[ref.Int()]
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Returns a slice containing the string representation of all valid temperature units
func (tu TemperatureUnit) UnitSlice() []string {
	return []string{"°C", "°F"}
}

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	return UnitString(tu)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Returns a slice containing the string representation of all valid speed units
func (su SpeedUnit) UnitSlice() []string {
	return []string{"km/h", "mph"}
}

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	return UnitString(su)
}

type Measurement interface {
	GetValue() int
	GetUnit() Unit
}

// Converts a Measurement to a string
func MeasurementString(measurement Measurement) string {
	return fmt.Sprintf("%d %v", measurement.GetValue(), measurement.GetUnit())
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Returns the degree
func (t Temperature) GetValue() int {
	return t.degree
}

// Returns the TemperatureUnit
func (t Temperature) GetUnit() Unit {
	return t.unit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return MeasurementString(t)
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Returns the magnitude
func (s Speed) GetValue() int {
	return s.magnitude
}

// Returns the SpeedUnit
func (s Speed) GetUnit() Unit {
	return s.unit
}

// Add a String method to Speed
func (s Speed) String() string {
	return MeasurementString(s)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %v, Wind %s at %v, %d%% Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}
