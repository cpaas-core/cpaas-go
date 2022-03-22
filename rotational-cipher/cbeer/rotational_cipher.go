package rotationalcipher

// RotMap is a rotation map for the cipher
type RotMap map[rune]rune

// CreateRotMap creates the rotation map for the cipher
func CreateRotMap(shiftKey int) RotMap {
	rotMap := RotMap{}
	lower := []rune("abcdefghijklmnopqrstuvwxyz")
	upper := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 0; i < len(lower); i++ {
		offset := (i + shiftKey) % len(lower)
		rotMap[lower[i]] = lower[offset]
		rotMap[upper[i]] = upper[offset]
	}
	return rotMap
}

// RotationalCipher encrypts the provided plan string using a rotational cypher
// based on the shiftKey value 
func RotationalCipher(plain string, shiftKey int) string {
	rotMap := CreateRotMap(shiftKey)
	cipherText := ""
	for _, letter := range plain {
		if val, ok := rotMap[letter]; ok {
			cipherText += string(val)
		} else {
			cipherText += string(letter)
		}
	}
	return cipherText
}
