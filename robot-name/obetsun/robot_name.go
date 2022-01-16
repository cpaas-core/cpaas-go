package robotname

import (
	"fmt"
)

type OutOfNamesError struct {
	message string
}

func (e *OutOfNamesError) Error() string {
	return e.message
}

// Define the Robot type here.
type Robot struct {
	name string
}

const LETTER_VARIATIONS = 26 * 26
const DIGITS_VARIATIONS = 10 * 10 * 10

var names = (func() map[string]bool {
	nms := make(map[string]bool)
	for r1 := 'A'; r1 <= 'Z'; r1++ {
		for r2 := 'A'; r2 <= 'Z'; r2++ {
			for i := 0; i <= 999; i++ {
				key := fmt.Sprintf("%c%c%003d", r1, r2, i)
				nms[key] = true
			}
		}
	}
	return nms
})()

func GenerateName() (string, error) {
	if len(names) <= 0 {
		return "", &OutOfNamesError{
			"Error: Out of names",
		}
	}
	var nm string
	for k, _ := range names {
		nm = k
		delete(names, k)
		break
	}
	return nm, nil
}

func (r *Robot) Name() (string, error) {
	var err error
	if len(r.name) != 5 {
		r.name, err = GenerateName()
	}
	return r.name, err
}

func (r *Robot) Reset() {
	r.name = ""
}
