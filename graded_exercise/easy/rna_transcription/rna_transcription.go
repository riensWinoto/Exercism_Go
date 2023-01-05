package main

import "fmt"

func ToRNA(dna string) string {
	dnaRune := []rune(dna)
	for idx, val := range dnaRune {
		switch val {
		case 'G':
			dnaRune[idx] = 'C'
		case 'C':
			dnaRune[idx] = 'G'
		case 'T':
			dnaRune[idx] = 'A'
		case 'A':
			dnaRune[idx] = 'U'
		}
	}
	return string(dnaRune)
}

func main() {
	fmt.Println(ToRNA("GCTA"))
}
