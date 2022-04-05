package rotationalcipher

import (
	"fmt"
	"strings"
)

const lcAlphabet = "abcdefghijklmnopqrstuvwxyz"
const ucAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"


func toPositiv(number int) int {
	if (number < 0 ) {
		return number * -1
	}
	return number
}

func Lookup(token string, key int) string {

	replace  := ""
	alphaLen := 26
	isUpperCase := false

	if key == alphaLen {
		key = 0
	}
	pos := strings.Index(lcAlphabet, token)
	if pos < 0 {
		pos = strings.Index(ucAlphabet, token) 
		if pos >= 0 {
			isUpperCase = true
		} else {
			return token
		}
	}
	rotate := pos + key
	if rotate > alphaLen -1 {
		rotate =  rotate - alphaLen

	} 
	if isUpperCase {
		replace = fmt.Sprintf("%c", ucAlphabet[toPositiv(rotate)])
	} else {
		replace = fmt.Sprintf("%c", lcAlphabet[toPositiv(rotate)])
	}
	return replace

}

func RotationalCipher(plain string, shiftKey int) string {

	shifted := ""
	for _, char := range(plain) {
		shifted += Lookup(string(char), shiftKey)
	}
	return shifted
}
