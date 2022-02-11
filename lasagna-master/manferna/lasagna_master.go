package lasagna

// averageTimePerLayer = 2 (minutes)
// averageTimeTotal in minutes
func PreparationTime(layers []string, averageTimePerLayer int) int {

	if averageTimePerLayer == 0 {
		averageTimePerLayer = 2
	}

	averageTimeTotal := 0

	averageTimeTotal = averageTimePerLayer * len(layers)

	return averageTimeTotal
}

// sauceForOneLayer 0.2 liters
// noodlesForOneLayer 50 grams
func Quantities(layers []string) (int, float64) {

	var sauceCount = 0
	var noodlesCount = 0
	const sauceForOneLayer = 0.2
	const noodlesForOneLayer = 50
	for _, layer := range layers {
		if layer == "sauce" {
			sauceCount = sauceCount + 1
		}
		if layer == "noodles" {
			noodlesCount = noodlesCount + 1
		}
	}

	noodleTotal := noodlesCount * noodlesForOneLayer
	sauceTotal := float64(float64(sauceCount) * sauceForOneLayer)
	return noodleTotal, sauceTotal
}

// friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
// myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}
func AddSecretIngredient(friendList []string, myList []string) {

	secretIngredient := friendList[len(friendList)-1]

	if myList[len(myList)-1] == "?" {
		myList[len(myList)-1] = secretIngredient
	}
}

//quantities in grams
//scaledQuantities in grams
func ScaleRecipe(quatities []float64, portions int) []float64 {

	var quatitiesScaled []float64
	var quantityScaled float64

	for _, quantity := range quatities {
		// the /2 is needed bc the initial recipe is for 2 serves!
		quantityScaled = (quantity / 2) * float64(portions)
		quatitiesScaled = append(quatitiesScaled, quantityScaled)
	}
	return quatitiesScaled
}
