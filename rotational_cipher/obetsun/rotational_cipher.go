package rotationalcipher

const ALPHA = "abcdefghijklmnopqrstuvwxyz"
const ALPHA_UP = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NOT_FOUND = 1024

func GetIndex(c rune, capital bool) int {
	alp := ALPHA
	if capital {
		alp = ALPHA_UP
	}

	for i, v := range alp {
		if v == c {
			return i
		}
	}
	return NOT_FOUND
}

func RotationalCipher(plain string, shiftKey int) string {
	var res string
	capital := false
	alp := ALPHA
	for _, v := range plain {
		if GetIndex(v, false) == GetIndex(v, true) {
			res += string(v)
		} else {
			if GetIndex(v, true) != NOT_FOUND {
				capital = true
				alp = ALPHA_UP
			}
			newIndex := GetIndex(v, capital) + shiftKey
			if newIndex > len(alp)-1 {
				newIndex -= len(alp)
			}
			newRune := alp[newIndex : newIndex+1]
			res += newRune
			capital = false
			alp = ALPHA
		}
	}
	return res
}
