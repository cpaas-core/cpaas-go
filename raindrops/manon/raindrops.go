package raindrops

import "strconv"

var Drops = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(number int) string {
	var response string
	var potentialFactor = 1
	for dropCounter := 1; dropCounter <= len(Drops); dropCounter += 0 {
		if sound, exists := Drops[potentialFactor]; exists == true {
			if number%potentialFactor == 0 {
				response += sound
			}
			dropCounter++
		}
		potentialFactor++
	}
	if response == "" {
		response = strconv.Itoa(number)
	}
	return response
}
