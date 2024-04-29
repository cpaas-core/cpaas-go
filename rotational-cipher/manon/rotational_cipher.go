package rotationalcipher

import "unicode"

const alphabetSize = 26

func RotationalCipher(plain string, shiftKey int) string {
	var cipher []rune
	for _, letter := range plain {
		// No need to cipher
		if !unicode.IsLetter(letter) {
			cipher = append(cipher, letter)
			continue
		}
		var firstLetter rune
		if unicode.IsUpper(letter) {
			firstLetter = 'A'
		} else {
			firstLetter = 'a'
		}
		// Algorithm
		cipherLetter := (letter-firstLetter+int32(shiftKey))%alphabetSize + firstLetter
		cipher = append(cipher, cipherLetter)
	}
	return string(cipher)
}
