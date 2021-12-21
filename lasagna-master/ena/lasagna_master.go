package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	const averageTimePerLayer = 2
	if timePerLayer == 0 {
		timePerLayer = averageTimePerLayer
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	const noodlesPerLayer = 50 // grams
	const saucePerLayer = 0.2  // liters
	var noodlesLayers, sauceLayers int
	for _, entry := range layers {
		if entry == "sauce" {
			sauceLayers++
		} else if entry == "noodles" {
			noodlesLayers++
		}
	}
	return noodlesLayers * noodlesPerLayer, float64(sauceLayers) * saucePerLayer
}

func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(quantities []float64, scaleFactor int) []float64 {
	var scaledQuantities []float64
	for _, entry := range quantities {
		scaledQuantities = append(scaledQuantities, (entry/2.0)*float64(scaleFactor))
	}
	return scaledQuantities
}
