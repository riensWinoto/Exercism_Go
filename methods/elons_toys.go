package main

import "fmt"

var (
	myCarSpeed, myCarBatteryDrain, myTrackDistance int = 5, 2, 100
)

type Car struct {
	speed, batteryDrain, battery, distance int
}

func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.distance = c.speed
		c.battery -= c.batteryDrain
	}
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	if (c.battery/c.batteryDrain)*c.speed >= trackDistance {
		return true
	} else {
		return false
	}
}

func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

func main() {
	myCar := NewCar(myCarSpeed, myCarBatteryDrain)
	myCar.Drive()
	fmt.Println(myCar)
	fmt.Println(myCar.DisplayDistance())
	fmt.Println(myCar.DisplayBattery())
	fmt.Println(myCar.CanFinish(myTrackDistance))
}
