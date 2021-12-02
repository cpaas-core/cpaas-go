package lasagna

// PreparationTime returns the estimate for the total preparation time
// as an int based on the number of layers.
func PreparationTime(layers []string, prepTimePerLayer int) int {
  if prepTimePerLayer < 1 {
    prepTimePerLayer = 2
  }
  return prepTimePerLayer * len(layers)
}

// Quantities determines the quantity of noodles and sauce needed to make
// your meal. The result should be returned as two values of noodles as
// an int and sauce as a float64.
func Quantities(layers []string) (int, float64) {
  var qtyNoodles int = 0
  var qtySauce float64 = 0.0
  for _, layer := range layers {
    if layer == "noodles" {
      qtyNoodles = qtyNoodles + 50
    } else if layer == "sauce" {
      qtySauce = qtySauce + 0.2
    }
  }
  return qtyNoodles, qtySauce
}

// AddSecretIngredient generates a new slice and adds the last item from
// your friend's list to the end of your list.
func AddSecretIngredient(friendsList, myList []string) []string {
  newList := make([]string, len(myList))
  copy(newList, myList)
  if len(friendsList) > 0 {
    newList = append(newList, friendsList[len(friendsList)-1])
  }
  return newList
}

// ScaleRecipe calculates the amounts for different numbers of portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
  var scalePortion float64 = float64(portions) / 2.0
  newAmounts := make([]float64, 0)
  for _, amount := range amounts {
    newAmounts = append(newAmounts, amount * scalePortion)
  }
  return newAmounts
}
