package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	var randomNumber int
	for {
		randomNumber = rand.Intn(21)
		if !(randomNumber == 0) {
			break
		}
	}
	return randomNumber
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	randomInRange := rand.Float64() * 12.0
	return randomInRange

}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	var animalsSuffle []string

	animalsSuffle = animals
	rand.Shuffle(len(animalsSuffle), func(i, j int) {
		animalsSuffle[i], animalsSuffle[j] = animalsSuffle[j], animalsSuffle[i]
	})
	return animalsSuffle
}
