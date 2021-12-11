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

	for _,layer := range layers{
		switch layer {
			case "noodles" : noodles++
			case "sauce" : sauces++ 
		}
	}
	noodles = noodles*50
	sauce = float64(sauces)*0.2
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList = append(myList, friendsList[len(friendsList)-1])
	return myList
}


func ScaleRecipe(amounts []float64, portions int) []float64 {
	res := make([]float64,0,len(amounts))
	scaleFactor := float64(portions) 
	for _,v  := range amounts{
		res = append(res, v/2*scaleFactor)
	}
	return res
}
