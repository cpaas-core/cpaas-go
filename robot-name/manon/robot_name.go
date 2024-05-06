package robotname

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

const (
	letters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers    = "0123456789"
	numLetters = 2
	numNumbers = 3
)

var (
	cache           = rand.Int63()
	maxCombinations = int(math.Pow(float64(len(letters)), numLetters) * math.Pow(float64(len(numbers)), numNumbers))
	usedNames       = make(map[string]bool)
)

// Robot struct
type Robot struct {
	name string
}

// Name set the name of the Robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(usedNames) == maxCombinations {
		return "", errors.New("No more names available")
	}
	name := getRandomName()
	for _, usedName := usedNames[name]; usedName; {
		name = getRandomName()
		_, usedName = usedNames[name]
	}
	r.name = name
	usedNames[name] = true
	return r.name, nil
}

// Reset remove the name of the Robot
func (r *Robot) Reset() {
	r.name = ""
}

func getRandomName() string {
	sb := strings.Builder{}
	sb.Grow(numLetters + numNumbers)

	// Fill random letters
	fillRandomChars(&sb, letters, 2)
	// Fill random numbers
	fillRandomChars(&sb, numbers, 3)

	return sb.String()
}

func fillRandomChars(sb *strings.Builder, charSet string, numChars int) {
	// Number of bits needed to represent the maximum index of the charset
	idxBits := len(strconv.FormatInt(int64(len(charSet)-1), 2))
	// All 1-bits, as many as idxBits
	idxMask := 1<<idxBits - 1

	if cache == 0 {
		cache = rand.Int63()
	}
	for ; numChars > 0; numChars-- {
		// The probability is not the same for all numbers when using the reminder
		idx := (cache & int64(idxMask)) % int64(len(charSet))
		sb.WriteByte(charSet[idx])
		cache >>= idxBits
	}
}
