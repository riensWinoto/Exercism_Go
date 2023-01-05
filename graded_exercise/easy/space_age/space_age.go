package main

import "fmt"

const earthSec float64 = 31557600

var (
	earth   Planet  = Planet("Earth")
	seconds float64 = 1000000000
)

type Planet string

func Age(seconds float64, planet Planet) float64 {
	planetMap := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	switch planet {
	case "Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune":
		return (seconds / earthSec) * (planetMap["Earth"] / planetMap[planet])
	default:
		return -1.000000
	}
}

func main() {
	fmt.Println(Age(seconds, earth))
}
