// Manage robot factory settings.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot has a name
type Robot struct {
	name string
}

// All the existing robot names so we can check for unique names
var robots = make(map[string]Robot)

// Name returns a robot's name in the format of two uppercase letters
// followed by three digits. If the robot doesn't have a name already,
// a random name is generated. Names must be random (no predictable pattern)
// and unique (no duplicate names).
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		newName, error := GenerateUniqueName()
		if error != nil {
			return "", error
		}
		r.name = newName
		robots[newName] = *r
	}
	return r.name, nil
}

// Reset clears the name of a robot so the next time a name is requested a
// new random name is generated.
func (r *Robot) Reset() {
	r.name = ""
}

// Maximum number of unique name possible (26x26x1000)
var MAX_NAMES = 676000
var src = rand.NewSource(time.Now().UnixNano())

// GenerateUniqueName creates a name and makes sure it isn't a duplicate.
func GenerateUniqueName() (string, error) {
	if len(robots) == MAX_NAMES {
		return "", errors.New("No more unique robot names available")
	}
	letter1 := rand.Intn(26)
	letter2 := rand.Intn(26)
	robotNumber := rand.Intn(1000)

	// Time to cheat. Rather than guess randomly, we create a unique
	// starting point and increment until we find a unique name
	// Since CreateName uses modulo to get the proper letters and numbers
	// we can just increment here.
	for firstOffset := 0; firstOffset < 26; firstOffset++ {
		for secondOffset := 0; secondOffset < 26; secondOffset++ {
			for numberOffset := 0; numberOffset < 1000; numberOffset++ {
				newName := CreateName(
					letter1+firstOffset,
					letter2+secondOffset,
					robotNumber+numberOffset)
				if _, ok := robots[newName]; !ok {
					return newName, nil
				}
			}
		}
	}

	// Should never get here, the initial check should catch this. But just in case...
	return "", errors.New("No really, there are no more unique robot names available!")
}

// CreateName given three ints, generate a robot name
// in the form AA123
func CreateName(letter1, letter2, robotNumber int) string {
	return fmt.Sprintf("%c%c%03d",
		'A'+rune(letter1%26),
		'A'+rune(letter2%26),
		robotNumber%1000)
}
