package main

import "fmt"

const successRate = 90

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	calculateWorkingCarsPerHour := float64(productionRate) * (successRate / 100)
	return float64(calculateWorkingCarsPerHour)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	temp := CalculateWorkingCarsPerHour(productionRate, successRate)
	calculateWorkingCarsPerMinute := temp / 60.0
	return int(calculateWorkingCarsPerMinute)
}

func CalculateCost(carsCount int) uint {
	perGroupTen := carsCount / 10
	perIdv := carsCount % 10
	calculateCost := (perGroupTen * 95000) + (perIdv * 10000)
	return uint(calculateCost)
}

func main() {
	workingCarsPerHour := CalculateWorkingCarsPerHour(1547, successRate)
	workingCarsPerMinute := CalculateWorkingCarsPerMinute(1105, successRate)
	cost := CalculateCost(37)
	fmt.Println("Working cars per hour ", workingCarsPerHour)
	fmt.Println("Working cars per minute ", workingCarsPerMinute)
	fmt.Println("Cost ", cost)
}
