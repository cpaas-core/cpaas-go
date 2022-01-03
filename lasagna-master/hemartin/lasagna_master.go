package lasagna

const noodleQuantityPerLayer int = 50
const sauceQuantityPerLayer float64 = 0.2
const originalQuantityPortions float64 = 2.0

func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime <= 0 {
		layerPrepTime = 2
	}
	return layerPrepTime * len(layers)

}

func Quantities(layers []string) (int, float64) {
	return QuantitiesFast(layers)
}

// This was the first implementation
func QuantitiesSlow(layers []string) (int, float64) {
	count := map[string]int{
		"noodles": 0,
		"sauce":   0,
	}

	for _, layer := range layers {
		if layer == "noodles" || layer == "sauce" {
			count[layer]++
		}

	}

	return count["noodles"] * noodleQuantityPerLayer, float64(count["sauce"]) * sauceQuantityPerLayer
}

// Performant implementation
func QuantitiesFast(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += noodleQuantityPerLayer
		case "sauce":
			sauce += sauceQuantityPerLayer
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendList, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	scaleFactor := float64(portions) / originalQuantityPortions

	for index, quantity := range quantities {
		scaledQuantities[index] = quantity * scaleFactor
	}

	return scaledQuantities
}
