package cards

func indexInBound(slice []int, index int) bool {
	return index >= 0 && index < len(slice)
}

func GetItem(slice []int, index int) (int, bool) {
	if indexInBound(slice, index) {
		return slice[index], true
	}
	return 0, false
}

func SetItem(slice []int, index, value int) []int {
	if indexInBound(slice, index) {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}

func PrefilledSlice(value, length int) []int {
	var slice []int
	for i := 0; i < length; i++ {
		slice = append(slice, value)
	}
	return slice
}

func RemoveItem(slice []int, index int) []int {
	if indexInBound(slice, index) {
		if len(slice)-1 == index {
			slice = slice[:index]
		} else {
			slice = append(slice[:index], slice[index+1:]...)
		}
	}
	return slice
}
