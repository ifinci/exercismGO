package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avr int) int {
	if avr == 0 {
		avr = 2
	}
	return len(layers) * avr
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myLen := len(myList)
	friendLen := len(friendList)
	myList[myLen-1] = friendList[friendLen-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	if quantities == nil {
		return []float64{}
	}
	amounts := make([]float64, len(quantities))
	scaleFactor := float64(portions) / 2.0
	for i, v := range quantities {
		amounts[i] = v * scaleFactor
	}
	return amounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
