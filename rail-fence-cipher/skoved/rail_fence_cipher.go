package railfence

import (
	"strings"
)

// takes the length of a string and the number of rails for a rail fence cipher
// and writes the sequence of rail numbers that each letter of the message will
// be written to a channel
func railOrderGenerator(length, rails int, ch chan<- int) {
	for i, curRail, inc := 0, 0, 1; i < length; i, curRail = i+1, curRail+inc {
		ch <- curRail

		nextRail := curRail + inc
		// check if the increment direction needs to change
		if nextRail >= rails || nextRail < 0 {
			inc *= -1
		}
	}
}

// takes a string and the number for rails for a rail cipher and returns the
// string encoded using a rail cipher
func Encode(message string, rails int) string {
	msgLen := len(message)
	railGenerator := make(chan int, msgLen)
	matrix := make([][]rune, rails)

	go railOrderGenerator(msgLen, rails, railGenerator)
	for _, char := range message {
		curRail := <-railGenerator
		matrix[curRail] = append(matrix[curRail], char)
	}

	var builder strings.Builder
	builder.Grow(msgLen)
	for rail := 0; rail < rails; rail++ {
		for _, char := range matrix[rail] {
			_, err := builder.WriteRune(char)
			if err != nil {
				panic("could not write rune to string builder")
			}
		}
	}
	return builder.String()
}

// takes a rail cipher encoded string and the number of rails and returns the
// decode message
func Decode(message string, rails int) string {
	// determine the number of characters that go on each rail
	msgLen := len(message)
	charsPerRail := make([]int, rails)
	for i, curRail, inc := 0, 0, 1; i < msgLen; i, curRail = i+1, curRail+inc {
		charsPerRail[curRail]++

		nextRail := curRail + inc
		// check if the increment direction needs to change
		if nextRail >= rails || nextRail < 0 {
			inc *= -1
		}
	}
	matrix := make([][]rune, rails)
	charsInRail := 0
	curRail := 0

	// write the encoded message onto the rails
	for _, char := range message {
		if charsInRail >= charsPerRail[curRail] {
			curRail++
			charsInRail = 0

			if curRail >= rails {
				panic("the number of rails is too low")
			}
		}

		matrix[curRail] = append(matrix[curRail], char)
		charsInRail++
	}

	// extract the decoded message from the rails
	railGenerator := make(chan int, msgLen)
	var builder strings.Builder
	builder.Grow(msgLen)
	go railOrderGenerator(msgLen, rails, railGenerator)
	for i := 0; i < msgLen; i++ {
		curRail := <-railGenerator
		rail := matrix[curRail]
		var char rune

		if len(rail) > 0 {
			// get the first character from the rail and remove it
			// from the slice
			char, matrix[curRail] = rail[0], rail[1:]
		} else {
			panic("invalid message")
		}

		_, err := builder.WriteRune(char)
		if err != nil {
			panic("Could not write byte to string builder")
		}
	}

	return builder.String()
}
