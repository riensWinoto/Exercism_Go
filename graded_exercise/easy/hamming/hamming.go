package main

import "fmt"

var (
	dnaA string = "GAGCCTACTAACGGGAT"
	dnaB string = "CATCGTAATGACGGCCT"
)

func Distance(a, b string) (int, error) {
	var diffNum int = 0
	if len(a) != len(b) {
		return -1, fmt.Errorf("MUST SAME LENGTH")
	}

	for idx := range a {
		if a[idx] != b[idx] {
			diffNum += 1
		}
	}
	return diffNum, nil
}

func main() {
	fmt.Println(Distance(dnaA, dnaB))
}
