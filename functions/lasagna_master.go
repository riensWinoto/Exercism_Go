package main

import "fmt"

var (
	layers               []string  = []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	averageTime, portion int       = 3, 4
	friendRecipe         []string  = []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myRecipe             []string  = []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	quantityRecipe       []float64 = []float64{1.2, 3.6, 10.5}
)

// PreparationTime calculate time needed for preparation
func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}
	return len(layers) * averageTime
}

// Quantities calculate how much gram of noodles and liter of sauce
func Quantities(layers []string) (noodlesGrams int, sauceLiter float64) {
	for _, val := range layers {
		if val == "noodles" {
			noodlesGrams += 50
		}
		if val == "sauce" {
			sauceLiter += 0.2
		}
	}
	return
}

// AddSecretIngredient modify recipe at the last index
func AddSecretIngredient(friendList []string, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
	fmt.Println(myList)
}

// ScaleRecipe calculate time needed to cook desired portion
func ScaleRecipe(oriQuantityRecipe []float64, portion int) []float64 {
	modQuantityRecipe := make([]float64, len(oriQuantityRecipe))

	for idx, val := range oriQuantityRecipe {
		modQuantityRecipe[idx] = (val / 2) * float64(portion)
	}
	return modQuantityRecipe
}

func main() {
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(Quantities(layers))
	AddSecretIngredient(friendRecipe, myRecipe)
	fmt.Println(ScaleRecipe(quantityRecipe, portion))
}
