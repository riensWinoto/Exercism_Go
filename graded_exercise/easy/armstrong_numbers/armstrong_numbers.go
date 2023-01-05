package main

import (
	"fmt"
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	temp := 0
	convertToStr := strconv.Itoa(n)
	countLength := len(convertToStr)
	for i := 0; i <= countLength-1; i++ {
		convertToInt, _ := strconv.Atoi(string(convertToStr[i]))
		temp += int(math.Pow(float64(convertToInt), float64(countLength)))
	}
	if temp == n {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsNumber(153))
}
