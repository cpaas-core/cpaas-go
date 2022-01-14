package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberUpperLimit = 999

var robotNumber = 0
var robotNames = func() []string {
	names := []string{}
	for _, firstLetter := range letters {
		for _, secondLetter := range letters {
			for number := 0; number <= numberUpperLimit; number++ {
				name := fmt.Sprintf("%c%c%03d", firstLetter, secondLetter, number)
				names = append(names, name)
			}
		}

	}
	return names
}()

func init() {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(robotNames), func(i, j int) { robotNames[i], robotNames[j] = robotNames[j], robotNames[i] })
}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if robotNumber >= len(robotNames) {
		return "", fmt.Errorf("robot names depleted")
	}

	r.name = robotNames[robotNumber]
	robotNumber++

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
