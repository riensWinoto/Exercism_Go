package main

import "fmt"

var year = 1990

func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	}
	return false
}

func main() {
	fmt.Println(IsLeapYear(year))
}
