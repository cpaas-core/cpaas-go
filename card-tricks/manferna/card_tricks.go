package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {

	ok := false
	card := 0

	if (index > len(slice)-1) || (index < 0) {
		return card, ok
	} else {
		ok = true
		card = slice[index]
	}
	return card, ok
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

	if (index > len(slice)-1) || (index < 0) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {

	var slice []int
	var i = 0

	if length > 0 {
		for {
			slice = append(slice, value)
			i = i + 1
			if i == length {
				break
			}
		}
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {

	var sliceOutOfIndex = slice

	if (index < 0) || (index > len(slice)-1) {
		return sliceOutOfIndex
	} else {
		var sliceBeforeIndex = slice[:index]
		var sliceAfterIndex = slice[(index + 1):]

		slice = append(sliceBeforeIndex, sliceAfterIndex...)
		return slice
	}
}
