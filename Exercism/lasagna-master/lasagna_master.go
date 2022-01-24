package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}

		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	secret := friendList[len(friendList)-1]
	myList[len(myList)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var slice = make([]float64, len(amounts))

	for i := 0; i < len(amounts); i++ {
		slice[i] = amounts[i] * float64(portions) / 2
	}

	return slice
}
