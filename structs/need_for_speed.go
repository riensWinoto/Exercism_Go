package main

import "fmt"

var (
	speedx        int = 5
	batteryDrainx int = 2
	distancex     int = 800
)

type Car struct {
	battery, batteryDrain, speed, distance int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	calculateBattery := car.battery - car.batteryDrain
	calculateDistance := car.distance + car.speed
	if car.battery < car.batteryDrain {
		return car
	}

	return Car{
		battery:      calculateBattery,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     calculateDistance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if (car.battery/car.batteryDrain)*car.speed >= track.distance {
		return true
	}
	return false
}

func main() {
	blackzero := NewCar(speedx, batteryDrainx)
	blackzero = Drive(blackzero)
	blacktrack := NewTrack(distancex)
	fmt.Println(blackzero)
	fmt.Println(blacktrack)
	fmt.Println(CanFinish(blackzero, blacktrack))
}
