package lasagna

//PreparationTime calculates the required preparation time based on
//quantity of layers in the layers array and average preparation time of a layer.
func PreparationTime(layers []string, averageTime int) int {
	if averageTime <= 0 {
		return len(layers) * 2
	}
	return len(layers) * averageTime
}

//Quantities returns the required quantities of noodles and sauce based
//on the layers provided.
func Quantities(layers []string) (noodles int, sauce float64) {
	var sauceAmount float64
	var noodleAmount int

	for _, layer := range layers {
		if layer == "sauce" {
			sauceAmount += 0.2
		}
		if layer == "noodles" {
			noodleAmount += 50
		}
	}
	return noodleAmount, sauceAmount
}

//AddSecretIngredient determines the secret ingredient from the
//friend's ingredient list and adds it to your own.
func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

//ScaleRecipe scales the given quantities for the desired number of
//portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled []float64
	for _, quantity := range quantities {
		scaled = append(scaled, quantity*float64(portions)/2)
	}
	return scaled
}
