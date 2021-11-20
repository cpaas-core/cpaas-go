package lasagna

func PreparationTime(layers []string, timePerLayer int)int {

	if timePerLayer == 0 {
		timePerLayer = 2
	}

	return len(layers)*timePerLayer

}

func Quantities(layers []string) (int,float64) {
	var noodles, sauces int
	var sauce float64

	for _,v := range layers{
		if v == "noodles" {
			noodles++
		} else if v == "sauce" {
			sauces++ 
		}
	}
	noodles = noodles*50
	sauce = float64(sauces)*0.2
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	var res []string = myList[:]
	res = append(res, friendsList[len(friendsList)-1])
	return res
}


func ScaleRecipe(amounts []float64, portions int) []float64 {
	res := make([]float64,0,len(amounts))
	for _,v  := range amounts{
		res = append(res, v/2*float64(portions))
	}
	return res
}

