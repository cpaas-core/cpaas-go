package robotname

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

var usedNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lettersPartLen = 2
	const digits = "0123456789"
	const digitsPartLen = 3

	if r.name != "" {
		return r.name, nil
	}

	allPossibleNames := int(math.Pow(float64(len(letters)), lettersPartLen) * math.Pow(float64(len(digits)), digitsPartLen))
	if len(usedNames) == allPossibleNames {
		return "", errors.New("All names taken")
	}

	for {
		name := GenRandomString(letters, lettersPartLen) + GenRandomString(digits, digitsPartLen)
		if _, exists := usedNames[name]; !exists {
			usedNames[name] = true
			r.name = name
			break
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func GenRandomString(pool string, length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}
