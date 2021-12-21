package cards

// isValidIndex determines if index is a valid position in slice
func isValidIndex(slice []int, index int) bool {
	return index >= 0 && index < len(slice)
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if isValidIndex(slice, index) {
		return slice[index], true
	}

	return 0, false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if isValidIndex(slice, index) {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value int, length int) (newSlice []int) {
	// This step of the excersize didn't give insight into handling negative length, but the test
	// case for negative length looks for a slice with no elements
	for i := 0; i < length; i++ {
		newSlice = append(newSlice, value)
	}
	return
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if isValidIndex(slice, index) {
		// thanks to 'Hymns For Disco' and 'T. Claverie' on stackoverflow for this cool
		// oneliner https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		return append(slice[:index], slice[index+1:]...)
	}

	return slice
}
