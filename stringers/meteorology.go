package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tempUnit TemperatureUnit) String() string {
	tempSliceStr := []string{"°C", "°F"}
	return tempSliceStr[tempUnit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {
	return fmt.Sprintf("%v %v", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string {
	speedSliceStr := []string{"km/h", "mph"}
	return speedSliceStr[speedUnit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (spd Speed) String() string {
	return fmt.Sprintf("%v %v", spd.magnitude, spd.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteorData MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		meteorData.location, meteorData.temperature, meteorData.windDirection,
		meteorData.windSpeed, meteorData.humidity)
}

func main() {
	celsiusTemp := Temperature{degree: 21, unit: Celsius}
	milesSpeed := Speed{magnitude: 100, unit: MilesPerHour}
	meteorDat := MeteorologyData{location: "Vrindavan",
		temperature:   Temperature{degree: 10, unit: Celsius},
		windDirection: "North West",
		windSpeed:     Speed{magnitude: 5, unit: KmPerHour},
		humidity:      88,
	}
	fmt.Println(Celsius)
	fmt.Println(celsiusTemp)
	fmt.Println(KmPerHour)
	fmt.Println(milesSpeed)
	fmt.Println(meteorDat)
}
