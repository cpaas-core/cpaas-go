package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {
	fmt.Println(message, rails)
	builders := make([]strings.Builder, rails)

	indexes := createEncodeIndexes(len(message), rails)
	fmt.Println(indexes)
	for letterPosition, letter := range message {
		railIndex := indexes[letterPosition]
		builders[railIndex].WriteRune(letter)
	}

	var finalString strings.Builder
	for _, builder := range builders {
		finalString.WriteString(builder.String())
	}
	return finalString.String()
}

// Tried, but no time left
func Decode(message string, rails int) string {
	return ""
}

func createEncodeIndexes(length int, totalRails int) []int {
	indexes := make([]int, length)

	increment := 1
	index := 0
	for i := 0; i < length; i++ {
		if index == 0 {
			increment = 1
		} else if index == totalRails-1 {
			increment = -1
		}

		indexes[i] = index
		index += increment
	}

	return indexes
}
