package main

import "fmt"

var (
	favCards   []int
	myCards    []int = []int{1, 2, 4, 1}
	myIndex    int   = 2
	myNewCards int   = 6
)

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favCards = []int{2, 6, 9}

	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {

	if index >= len(slice) || index < 0 {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		return append(slice, value)
	}

	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}

	values = append(values, slice...)
	return values
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var modifiedSlice []int

	if index >= len(slice) || index < 0 {
		return slice
	}

	modifiedSlice = slice[:index]
	slice = slice[index+1:]
	return append(modifiedSlice, slice...)
}

func main() {
	fmt.Println(FavoriteCards())
	fmt.Println(GetItem(myCards, myIndex))
	fmt.Println(SetItem(myCards, myIndex, myNewCards))
	fmt.Println(PrependItems(myCards, myIndex))
	fmt.Println(RemoveItem(myCards, myIndex))
}
