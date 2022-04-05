package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var encoded strings.Builder
	for _, char := range plain {
		var newChar rune

		if !unicode.IsLetter(char) {
			newChar = char
		} else if unicode.IsLower(char) {
			// Apply the shift and calculate the off limit it is from 'a'
			// For example, 'z'%'a' = 25, '{'%'a' = 26
			newShift := (char + rune(shiftKey)) % 'a'
			// Apply the offset from 'a' reseting it to 0 if 26
			newChar = 'a' + rune(newShift%26)
		} else if unicode.IsUpper(char) {
			newShift := (char + rune(shiftKey)) % 'A'
			newChar = 'A' + rune(newShift%26)
		}

		encoded.WriteRune(newChar)
	}

	return encoded.String()
}
