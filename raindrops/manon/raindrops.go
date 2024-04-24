package raindrops

import "strconv"

var drops = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var response string
	for _, drop := range drops {
		if number%drop.factor == 0 {
			response += drop.sound
		}
	}
	if response == "" {
		return strconv.Itoa(number)
	}
	return response
}
