package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	const noodleGrams = 50
	const sauceLiters = 0.2

	var noodlesLayers = 0
	var sauceLayers = 0.0

	for _, v := range layers {
		if v == "noodles" {
			noodlesLayers++
		}

		if v == "sauce" {
			sauceLayers++
		}
	}

	return noodlesLayers * noodleGrams, sauceLayers * sauceLiters
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuanities []float64
	for _, v := range quantities {
		scaledQuanities = append(scaledQuanities, (v/2)*float64(portions))
	}

	return scaledQuanities
}
