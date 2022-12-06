package main

import (
	"fmt"
	"strconv"
)

var coolNumber int = 888

func Convert(number int) string {
	var dropper string
	if number%3 == 0 {
		dropper += "Pling"
	}

	if number%5 == 0 {
		dropper += "Plang"
	}

	if number%7 == 0 {
		dropper += "Plong"
	}

	if dropper == "" {
		dropper = strconv.Itoa(number)
	}

	return dropper
}

func main() {
	fmt.Print(Convert(coolNumber))
}
