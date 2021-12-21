package cards

func isOutOfBounds(slice []int, index int) bool {
	return index < 0 || index >= len(slice)
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// at the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if isOutOfBounds(slice, index) {
		return 0, false
	}

	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is appended.
func SetItem(slice []int, index, value int) []int {
	if isOutOfBounds(slice, index) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) (slice []int) {
	for i := 0; i < length; i++ {
		slice = append(slice, value)
	}

	return
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if !isOutOfBounds(slice, index) {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}
