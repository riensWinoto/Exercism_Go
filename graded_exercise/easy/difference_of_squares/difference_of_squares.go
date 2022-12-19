package main

import (
	"fmt"
)

var number int = 10

func SquareOfSum(n int) int {
	var sumholder, temp int
	for i := 1; i <= n; i++ {
		sumholder += i
	}
	temp = sumholder * sumholder
	return temp
}

func SumOfSquares(n int) int {
	var temp int
	for i := 1; i <= n; i++ {
		multiplier := i * i
		temp += multiplier
	}
	return temp
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	fmt.Println(SquareOfSum(number))
	fmt.Println(SumOfSquares(number))
	fmt.Println(Difference(number))
}
