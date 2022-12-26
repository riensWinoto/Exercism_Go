package main

import "fmt"

type (
	File       []bool
	Chessboard map[string]File
)

var (
	myChessboard Chessboard = Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, true, false, false, false, false, false},
		"D": File{false, false, false, false, false, false, false, false},
		"E": File{false, false, false, false, false, true, false, true},
		"F": File{false, false, false, false, false, false, false, false},
		"G": File{false, false, false, true, false, false, false, false},
		"H": File{true, true, true, true, true, true, false, true},
	}
	myFile string = "A"
)

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileCounter := 0
	for _, val := range cb[file] {
		if val == true {
			fileCounter += 1
		}
	}
	return fileCounter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	rankCounter := 0
	for _, valFromCb := range cb {
		for idx, valInCb := range valFromCb {
			if idx == rank-1 {
				if valInCb == true {
					rankCounter += 1
				}
			}
		}
	}
	return rankCounter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for _, val := range cb {
		counter += len(val)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, valFromCb := range cb {
		for _, valInCb := range valFromCb {
			if valInCb == true {
				counter += 1
			}
		}
	}
	return counter
}

func main() {
	fmt.Println(CountInFile(myChessboard, myFile))
	fmt.Println(CountInRank(myChessboard, 2))
	fmt.Println(CountAll(myChessboard))
	fmt.Println(CountOccupied(myChessboard))
}
