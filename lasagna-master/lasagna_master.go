package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, average int) int {
	if average <= 0 {
		average = 2
	}

	return average * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	var noodlesInLayers int
	var sauceInLayers int

	for _,layer := range layers {
		if layer == "noodles" {
			noodlesInLayers += 1
		}

		if layer == "sauce" {
			sauceInLayers += 1
		}
	}

	noodles = noodlesInLayers * 50
	sauce = float64(sauceInLayers) * 0.2

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var response []float64

	for _,amount := range amounts {
		response = append(response, (amount / 2) * float64(portions))
	}

	return response
}
