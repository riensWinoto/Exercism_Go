package main

import (
	"fmt"
	"strings"
)

var (
	myWords string = "abcdefghijklmnopqrstuvwxyz"
)

func Score(word string) int {
	var myCounterScore int = 0

	for _, v := range word {
		valueUpper := strings.ToUpper(string(v))
		switch valueUpper {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			myCounterScore += 1

		case "D", "G":
			myCounterScore += 2

		case "B", "C", "M", "P":
			myCounterScore += 3

		case "F", "H", "V", "W", "Y":
			myCounterScore += 4

		case "K":
			myCounterScore += 5

		case "J", "X":
			myCounterScore += 8

		case "Q", "Z":
			myCounterScore += 10
		}
	}

	return myCounterScore
}

func main() {
	fmt.Println(Score(myWords))
}
