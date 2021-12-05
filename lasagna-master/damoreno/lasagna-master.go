package lasagna

func PreparationTime(layers []string, avgPreparationTimePerLayer int) int {
	if avgPreparationTimePerLayer < 1 {
		avgPreparationTimePerLayer = 2
	}
	return len(layers) * avgPreparationTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles++
		case "sauce":
			sauce++
		}
	}

	return noodles * 50, sauce * 0.2
}

func AddSecretIngredient(friendRecipe, myRecipe []string) []string {
	return append(myRecipe, friendRecipe[len(friendRecipe)-1])
}

func ScaleRecipe(amounts []float64, portions int) (scaledRecipe []float64) {
	multiplier := float64(portions) / 2.0

	for _, amount := range amounts {
		scaledRecipe = append(scaledRecipe, amount*multiplier)
	}

	return
}
