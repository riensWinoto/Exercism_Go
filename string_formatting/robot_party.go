package main

import "fmt"

var (
	myName      string  = "Riens"
	myAge       int     = 1000
	myTable     int     = 7
	myNeighbor  string  = "GoPyth"
	myDirection string  = "on the south east"
	myDistance  float64 = 34.55889
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf(`Welcome to my party, %s!
You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.
You will be sitting next to %s.`, name, table, direction, distance, neighbor)
}

func main() {
	fmt.Println(Welcome(myName))
	fmt.Println(HappyBirthday(myName, myAge))
	fmt.Println(AssignTable(myName, myTable, myNeighbor, myDirection, myDistance))
}
