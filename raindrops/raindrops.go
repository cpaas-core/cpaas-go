package raindrops

import "strconv"

// Convert a number into the sound of raindrops
func Convert(number int) string {
	sound := ""
	if (number % 3) == 0 {
		sound = "Pling"
	}
	if (number % 5) == 0 {
		sound += "Plang"
	}
	if (number % 7) == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = strconv.Itoa(number)
	}
	return sound
}
