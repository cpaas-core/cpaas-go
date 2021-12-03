package lasagna


// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
    if time != 0{
        return len(layers) * time
    }
    return len(layers) * 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodles, sauce := 0, 0.0
    for _, ingredient := range layers {
		if ingredient == "noodles"{
            noodles += 50
        }
    	if ingredient == "sauce"{
            sauce += 0.2
        }
	}
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string{
    return append(myList, friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64{
   	var amounts []float64
    share := float64(portions)/2.0
    for _, quantity := range quantities {
        amounts = append(amounts, quantity * share)
    }
	return amounts
    
    
}
