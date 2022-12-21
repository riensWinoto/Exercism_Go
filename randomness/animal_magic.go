package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	newSeedValue := time.Now().UnixNano()
	rand.Seed(newSeedValue)
	fmt.Println(newSeedValue)
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	rollMe := rand.Intn(20)
	if rollMe == 0 {
		rollMe += 1
	}
	return rollMe
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animalsSlice := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animalsSlice), func(i, j int) { animalsSlice[i], animalsSlice[j] = animalsSlice[j], animalsSlice[i] })
	return animalsSlice
}

func main() {
	SeedWithTime()
	fmt.Println(RollADie())
	fmt.Println(GenerateWandEnergy())
	fmt.Println(ShuffleAnimals())
}
