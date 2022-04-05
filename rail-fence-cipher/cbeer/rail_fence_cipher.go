package railfence

type railsMap map[int]string

func Encode(message string, rails int) string {
	encodingRails := railsMap{}

	railNumber := 0
	railOffset := 1
	for _, letter := range message {
		encodingRails[railNumber] += string(letter)
		railNumber += railOffset
		if railNumber == (rails - 1) {
			railOffset = -1
		} else if railNumber == 0 {
			railOffset = 1
		}
	}

	encodedString := ""
	for i := 0; i < rails; i++ {
		encodedString += encodingRails[i]
	}
	return encodedString
}

func Decode(message string, rails int) string {
	// Figure out how many letters go into each rail
	decodingRailsCount := make(map[int]int)

	railNumber := 0
	railOffset := 1
	for range message {
		decodingRailsCount[railNumber] += 1
		railNumber += railOffset
		if railNumber == (rails - 1) {
			railOffset = -1
		} else if railNumber == 0 {
			railOffset = 1
		}
	}

	// Create slices for each rail
	decodingRails := railsMap{}
	startingChar := 0
	endingChar := 0
	for i := 0; i < rails; i++ {
		endingChar = startingChar + decodingRailsCount[i]
		decodingRails[i] = string(message[startingChar:endingChar])
		startingChar = endingChar
	}

	decodedMessage := ""
	railNumber = 0
	railOffset = 1
	for range message {
		runeString := []rune(decodingRails[railNumber])
		if len(runeString) > 1 {
			decodedMessage += string(runeString[0:1])
			decodingRails[railNumber] = string(runeString[1:])
		} else if len(runeString) == 1 {
			decodedMessage += string(runeString)
			decodingRails[railNumber] = ""
		}

		railNumber += railOffset
		if railNumber == (rails - 1) {
			railOffset = -1
		} else if railNumber == 0 {
			railOffset = 1
		}
	}

	return decodedMessage
}
