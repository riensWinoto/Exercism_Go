package main

import (
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	var myNum uint64
	if number <= 0 || number > 64 {
		return 0, fmt.Errorf("Insert proper value")
	}
	myNum = uint64(math.Pow(2, float64(number-1)))
	return myNum, nil

}

func Total() uint64 {
	var sumsHolder uint64
	for i := 1; i <= 64; i++ {
		temp, err := Square(i)
		if err != nil {
			fmt.Println(err)
		}
		sumsHolder += temp
	}
	return sumsHolder
}

func main() {
	fmt.Println(Square(16))
	fmt.Println(Total())
}
