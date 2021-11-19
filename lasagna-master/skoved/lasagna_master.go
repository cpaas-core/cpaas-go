package lasagna

const defaultLayerPrepTime = 2
const noodlesPerLayer int = 50
const saucePerLayer float64 = 0.2
const defaultPortions float64 = 2.0

func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime <= 0 {
		layerPrepTime = defaultLayerPrepTime
	}
	return len(layers) * layerPrepTime
}

func Quantities(layers []string) (int, float64) {
	totalNoodles := 0
	totalSauce := 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			totalNoodles += noodlesPerLayer
		case "sauce":
			totalSauce += saucePerLayer
		}
	}
	return totalNoodles, totalSauce
}

func AddSecretIngredient(friendList []string, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := float64(portions) / defaultPortions
	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = quantity * scale
	}
	return scaledQuantities
}
