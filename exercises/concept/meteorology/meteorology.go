// package meteorology
package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (unit TemperatureUnit) String() string {
	if unit == 0 {
		return "°C"
	}
	if unit == 1 {
		return "°F"
	}
	return "wtf is this unit? "
	//return [...]string{"°C", "°F"}[unit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {
	return fmt.Sprint(temp.degree, " ", temp.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	return [...]string{"km/h", "mph"}[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprint(s.magnitude, " ", s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %d %s, Wind %s at %d %s, %d%% Humidity",
		md.location, md.temperature.degree, md.temperature.unit, md.windDirection,
		md.windSpeed.magnitude, md.windSpeed.unit, md.humidity)
}


func main() {
	celUnit := Celsius
	fmt.Println(celUnit.String())
	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit: Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit: MilesPerHour,
		},
		humidity: 60,
	}
	
	fmt.Println(sfData.String())
	// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
	// fmt.Sprint(sfData) 
}
