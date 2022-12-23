package main

import "fmt"

var birdsPerDay []int = []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var summingBirds int
	for _, val := range birdsPerDay {
		summingBirds += val
	}
	return summingBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var (
		summingBirds int
		birdsPerWeek []int
	)
	switch week {
	case 1:
		birdsPerWeek = birdsPerDay[0:7]
	case 2:
		birdsPerWeek = birdsPerDay[7:14]
	default:
		birdsPerWeek = birdsPerDay[0:7]
	}
	for _, val := range birdsPerWeek {
		summingBirds += val
	}
	return summingBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx, val := range birdsPerDay {
		if idx == 0 || idx%2 == 0 {
			birdsPerDay[idx] = val + 1
		}
	}
	return birdsPerDay
}

func main() {
	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, 1))
	fmt.Println(FixBirdCountLog(birdsPerDay))
}
