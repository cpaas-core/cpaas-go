package rotationalcipher

import "unicode"

func rotateRune(c rune, shiftKey int) rune {
	var ordA int
	if unicode.IsUpper(c) {
		ordA = int('A')
	} else {
		ordA = int('a')
	}
	cIdx := int(c) - ordA
	shifted := (cIdx + shiftKey) % 26
	return rune(shifted + ordA)
}
func RotationalCipher(plain string, shiftKey int) string {
	res := []rune{}
	for _, c := range plain {
		if unicode.IsLetter(c) {
			c = rotateRune(c, shiftKey)
		}
		res = append(res, c)
	}
	return string(res)
}
