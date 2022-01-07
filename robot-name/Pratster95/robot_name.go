package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var nameUsed = map[string]bool{}

const max = 26 * 26 * 10 * 10 * 10

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(nameUsed) >= max {
		return "", fmt.Errorf("no names left")
	}
	for {
		name := getName()
		if !nameUsed[name] {
			nameUsed[name] = true
			r.name = name
			return r.name, nil
		}
	}
}

func (r *Robot) Reset() {
	r.name = ""
}
func getName() string {
	return fmt.Sprintf("%s%s%d", string(rand.Intn(26)+66), string(rand.Intn(26)+66), rand.Intn(1000))

}
