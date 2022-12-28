package main

import "fmt"

var (
	residentName   string            = "Riens"
	residentAge    int               = 100
	residentAdress map[string]string = map[string]string{"street": "Main St."}
)

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || r.Address == nil || len(r.Address) == 0 {
		return false
	}
	if r.Address["street"] == "" {
		return false
	}

	for key := range r.Address {
		if key != "street" {
			return false
		}
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	counter := 0
	for _, val := range residents {
		if val.HasRequiredInfo() {
			counter += 1
		}
	}
	return counter
}

func main() {
	newCrazyRich := NewResident(residentName, residentAge, residentAdress)
	newListCrazyRich := []*Resident{newCrazyRich}
	fmt.Println(newCrazyRich)
	fmt.Println(newCrazyRich.HasRequiredInfo())
	newCrazyRich.Delete()
	fmt.Printf("%s %d %s\n", newCrazyRich.Name, newCrazyRich.Age, newCrazyRich.Address)
	fmt.Println(Count(newListCrazyRich))
}
