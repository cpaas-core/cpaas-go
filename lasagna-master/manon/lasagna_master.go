package lasagna

import "errors"

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers)*timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0 
	for i := 0; i < len(layers); i++ {
		layer := layers[i]
		switch layer{
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendIngredients, ownIngredients []string) {
	secretIngredient := friendIngredients[len(friendIngredients) - 1]
	unknownIngredient := ownIngredients[len(ownIngredients) - 1]
	if unknownIngredient == "?" {
		ownIngredients[len(ownIngredients) - 1] = secretIngredient
	}else{
		errors.New("Expected unknown ingredient not found")
	}
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledQuantities := make([]float64, 0, len(quantities))
	for _,quantity := range quantities {
		scaledQuantities = append(scaledQuantities, quantity/2.0*float64(numPortions))
	}
	return scaledQuantities
}
