package main

const successRate = 90

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	calculateWorkingCarsPerHour := float64(productionRate) * (successRate / 100)
    return float64(calculateWorkingCarsPerHour)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	temp := CalculateWorkingCarsPerHour(productionRate, successRate)
    calculateWorkingCarsPerMinute := temp/60.0
    return int(calculateWorkingCarsPerMinute)
}

func CalculateCost(carsCount int) uint {
    perGroupTen := carsCount / 10
    perIdv := carsCount % 10
    calculateCost := (perGroupTen * 95000) + (perIdv * 10000) 
    return uint(calculateCost)
}

func main(){
    CalculateWorkingCarsPerHour(1547, successRate)
    CalculateWorkingCarsPerMinute(1105, successRate)
    CalculateCost(37)
}
