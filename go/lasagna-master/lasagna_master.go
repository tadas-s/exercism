package lasagnamaster

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	noodleLayers := 0
	sauceLayers := 0

	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayers++
		} else if layer == "sauce" {
			sauceLayers++
		}
	}

	return noodleLayers * 50, float64(sauceLayers) * 0.2
}

func AddSecretIngredient(friendsList, myList []string) {
	if len(myList) < 1 {
		return
	}

	if myList[len(myList)-1] == "?" {
		if len(friendsList) < 1 {
			return
		}

		myList[len(myList)-1] = friendsList[len(friendsList)-1]
	}
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))

	for i, quantity := range quantities {
		result[i] = quantity / 2 * float64(portions)
	}

	return result
}
