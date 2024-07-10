package railfence

import (
	"math"
	"strings"
)

func Encode(message string, rails int) string {
	return partialEncode(rails, "", []string{message})
}

// partialEncode allows a tail recursion for the encoding algorithm
func partialEncode(rails int, encoded string, messages []string) string {
	var rest []string
	if rails <= 1 {
		return encoded + strings.Join(messages, "")
	}
	interval := (rails-2)*2 + 2
	for _, message := range messages {
		for i := 0; i < len(message); i += interval {
			encoded += string(message[i])
			rest = append(rest, message[i+1:min(i+interval, len(message))])
		}
	}
	return partialEncode(rails-1, encoded, rest)
}

func Decode(message string, rails int) string {
	// Set a slice of strings that represent the rails
	var railData []string
	var pointer, railLength int
	interval := (rails-2)*2 + 2
	for i := 1; i <= rails; i++ {
		if i == 1 {
			railLength = int(math.Ceil(float64(len(message)) / float64(interval)))
		} else if i == rails {
			railLength = int(math.Floor(float64(len(message)) / float64(interval)))
		} else {
			railLength = int(math.Round(float64(len(message)) / float64(interval) * 2))
		}
		railData = append(railData, message[pointer:pointer+railLength])
		pointer += railLength
	}
	// Go through rails picking one rune of each rail
	increase := true
	rail := 0
	var decoded string
	for i := 0; i < len(message); i++ {
		decoded += string(railData[rail][0])
		railData[rail] = railData[rail][1:]
		if increase {
			rail++
		} else {
			rail--
		}
		if rail == rails-1 || rail == 0 {
			increase = !increase
		}
	}
	return decoded
}
