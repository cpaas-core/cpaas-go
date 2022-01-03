package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < 0 || len(slice) == 0 {
		return 0, false
	}

	if index > len(slice)-1 {
		return 0, false
	}

	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if _, ok := GetItem(slice, index); ok {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}

	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length <= 0 {
		return []int{}
	}

	n := 0
	slice := []int{}

	for n < length {
		slice = append(slice, value)
		n++
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if _, ok := GetItem(slice, index); ok {
		new := []int{}
		new = append(new, slice[:index]...)
		new = append(new, slice[index+1:]...)
		return new
	}
	return slice
}
