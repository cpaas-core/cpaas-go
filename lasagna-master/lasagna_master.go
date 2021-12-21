package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
	   timePerLayer = 2
   }
   return timePerLayer * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += .2
		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList) - 1])
}

func ScaleRecipe(quantities []float64, amounts int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
	 scaledQuantities[i] = quantity * float64(amounts) / 2
	}
	return scaledQuantities
}

 

