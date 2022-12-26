package main

import "fmt"

const (
	item string = "carrot"
	unit string = "dozen"
)

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitScore := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return unitScore
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	customerBill := make(map[string]int)
	return customerBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, isAvail := units[unit]
	if isAvail != true {
		return false
	} else {
		bill[item] += units[unit]
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemPiece, isAvailBill := bill[item]
	unitQty, isAvailUnit := units[unit]

	if isAvailUnit != true || isAvailBill != true || unitQty > itemPiece {
		return false
	} else if itemPiece-unitQty == 0 {
		delete(bill, item)
	} else {
		bill[item] -= unitQty
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemPiece, isAvailBill := bill[item]
	return itemPiece, isAvailBill
}

func main() {
	unitsFunc := Units()
	billFunc := NewBill()
	fmt.Println(unitsFunc)
	fmt.Println(billFunc)
	fmt.Println(AddItem(billFunc, unitsFunc, item, unit))
	fmt.Println(RemoveItem(billFunc, unitsFunc, item, unit))
	fmt.Println(GetItem(billFunc, item))
}
