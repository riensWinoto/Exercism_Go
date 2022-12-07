package main

import "fmt"

var (
	needLicenseVehicle string  = "truck"
	chooseVehicle1     string  = "Toyota Z"
	chooseVehicle2     string  = "Toyota X"
	originalPrice      float64 = 5000.0
	vehicleAge         float64 = 15.0
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var logix bool
	if kind == "car" || kind == "truck" {
		logix = true
	} else {
		logix = false
	}
	return logix
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	}
	return option1 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64
	if age < 3 {
		resellPrice = float64(originalPrice * 0.8)
	}
	if age >= 10 {
		resellPrice = float64(originalPrice * 0.5)
	}
	if age >= 3 && age < 10 {
		resellPrice = float64(originalPrice * 0.7)
	}
	return resellPrice
}

func main() {
	fmt.Println(NeedsLicense(needLicenseVehicle))
	fmt.Println(ChooseVehicle(chooseVehicle1, chooseVehicle2))
	fmt.Println(CalculateResellPrice(originalPrice, vehicleAge))
}
