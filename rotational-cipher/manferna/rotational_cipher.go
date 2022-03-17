package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var plainAlphabet = getAlphabet()
	var newAlphabet = generateAlphabet(shiftKey)
	var newText string

	for _, character := range plain {
		var stringifiedCharacter = string(character)
		// no need to cipher
		if !unicode.IsLetter(character) {
			newText = newText + stringifiedCharacter
			continue
		}
		// capital letters
		isUpper := unicode.IsUpper(character)
		if isUpper {
			stringifiedCharacter = strings.ToLower(stringifiedCharacter)
		}
		for index, letter := range plainAlphabet {
			if letter == stringifiedCharacter {
				if isUpper {
					newText = newText + strings.ToUpper(newAlphabet[index])
					break
				} else {
					newText = newText + newAlphabet[index]
					break
				}
			}
		}
	}
	return newText
}

func generateAlphabet(shiftKey int) []string {
	var alphabetRotated = make([]string, 26)
	var plainAlphabet = getAlphabet()
	alphabetRotated = append(plainAlphabet[shiftKey:], plainAlphabet[:shiftKey]...)

	return alphabetRotated
}

func getAlphabet() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
}

/*
func isSpace(character string) bool {
	if character == " " {
		return true
	}
	return false
}

func isNumber(character string) bool {
	_, err := strconv.Atoi(character)
	if err == nil {
		return true
	}
	return false
}

func isNumberOrSpace(character string) bool {
	return isNumber(character) || isSpace(character)
}
*/
