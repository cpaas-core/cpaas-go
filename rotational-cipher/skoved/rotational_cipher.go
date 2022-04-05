package rotationalcipher

import (
	"errors"
	"strings"
)

func isUpper(char rune) (bool, error) {
	if char >= 'A' && char <= 'Z' {
		return true, nil
	} else if char >= 'a' && char <= 'z' {
		return false, nil
	}
	return false, errors.New("Not a letter in the alphabet")
}

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	var rot strings.Builder
	rot.Grow(len(plain))
	for _, char := range plain {
		var begin, end rune

		upperCase, err := isUpper(char)
		if err != nil {
			rot.WriteRune(char)
			continue
		}

		if upperCase {
			begin, end = 'A', 'Z'
		} else {
			begin, end = 'a', 'z'
		}
		newChar := char + rune(shiftKey)
		if newChar > end {
			newChar = newChar - end + begin - rune(1)
		}
		rot.WriteRune(newChar)
	}

	return rot.String()
}
