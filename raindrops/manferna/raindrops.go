package raindrops

import (
	"strconv"
)

func Convert(number int) string {

	const threeAsFactor = "Pling"
	const fiveAsFactor = "Plang"
	const sevenAsFactor = "Plong"

	var raindropsResult = ""

	// potencial factor 3
	if number%3 == 0 {
		raindropsResult = raindropsResult + threeAsFactor
	}

	//potencial factor 5
	if number%5 == 0 {
		raindropsResult = raindropsResult + fiveAsFactor
	}

	//potencial factor 7
	if number%7 == 0 {
		raindropsResult = raindropsResult + sevenAsFactor
	}

	if raindropsResult == "" {
		numberToString := strconv.Itoa(number)
		return numberToString
	} else {
		return raindropsResult
	}

}
