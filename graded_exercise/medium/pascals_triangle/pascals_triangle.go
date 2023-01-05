package main

import "fmt"

func factorial(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	return sum
}

func Triangle(row int) [][]int {
	res := [][]int{}
	for i := 0; i < row; i++ {
		insideRow := make([]int, i+1)
		for j := 0; j <= i; j++ {
			top := factorial(i)
			bottom := factorial(j) * factorial(i-j)
			insideRow[j] = top / bottom
		}
		res = append(res, insideRow)
	}
	return res
}

func main() {
	fmt.Println(Triangle(7))
}
