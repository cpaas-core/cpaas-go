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

	noodlesLayers := 0
	sauceLayers := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesLayers++
		case "sauce":
			sauceLayers++
		}
	}

	return noodlesLayers * noodleGrams, sauceLayers * sauceLiters
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuanities []float64) {
	ratio := float64(portions) / 2

	for _, quantity := range quantities {
		scaledQuanities = append(scaledQuanities, ratio*quantity)
	}

	return scaledQuanities
}
