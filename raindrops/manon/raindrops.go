package raindrops

import "strconv"

func Convert(number int) string {
	var response string
	if number%3 == 0 {
		response += "Pling"
	}
	if number%5 == 0 {
		response += "Plang"
	}
	if number%7 == 0 {
		response += "Plong"
	}
	if response == "" {
		response = strconv.Itoa(number)
	}
	return response
}
