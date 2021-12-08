package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if (index >= 0) && (index < len(slice)) {
		return slice[index], true
	} else {
		return 0, false
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if (index >= 0) && (index < len(slice)) {
		slice[index] = value
		return slice
	} else {
		newSlice := make([]int, len(slice)+1)
		newSlice = append(slice, value)
		return newSlice
	}
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length > 0 {
		newSlice := make([]int, length)
		for i := 0; i < length; i++ {
			newSlice[i] = value
		}
		return newSlice
	} else {
		return nil
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	_, ok := GetItem(slice, index)
	if ok {
		newSlice := make([]int, 0)
		for i := 0; i < len(slice); i++ {
			if i != index {
				newSlice = append(newSlice, slice[i])
			}
		}
		return newSlice
	} else {
		return slice
	}
}
