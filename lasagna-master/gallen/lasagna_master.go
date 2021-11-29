package lasagna

func PreparationTime(layers []string, avgPreparationTime int) int {
    if (avgPreparationTime == 0) {
        avgPreparationTime = 2
    }
    return len(layers) * avgPreparationTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        if (layer == "noodles") {
            noodles += 50
        } else if (layer == "sauce") {
            sauce += 0.2
        }
    }
    return
}

func AddSecretIngredient(friendsList, myList []string) []string {
    return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
    var result []float64
    for _, v := range amounts {
        result = append(result, v / float64(2) * float64(portions))
    }
    return result
}
