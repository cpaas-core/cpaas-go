package chance

import (
	"math/rand"
	"strings"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
    t := time.Now()
    rand.Seed(t.UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	min := 1
	max := 20
	r := rand.Intn(max - min) +1
	return r
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	min := 0.0
	max := 12.0
	r := rand.Float64() * (max - min)
	return r
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	SeedWithTime()
	a := strings.Fields("ant beaver cat dog elephant fox giraffe hedgehog")
	rand.Shuffle(len(a), func (i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}
