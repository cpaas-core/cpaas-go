package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.


type SillyNephewError struct {
	err int 
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.err)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	if (cows < 0 ) {
		return 0.0, &SillyNephewError{err: cows}
	}

	avail, err := weightFodder.FodderAmount()
	if (err != nil) {
		if (errors.Is(err, ErrScaleMalfunction) && avail > 0) {
			avail = avail * 2
			err = nil
		} else if (errors.Is(err, ErrScaleMalfunction) && avail < 0) {
			avail = 0.0
			err = errors.New("negative fodder")
		} else {
			avail = 0.0
		}
	} else {
		if (avail < 0) {
			avail = 0.0
			err = errors.New("negative fodder")
		}
	}
	q := avail / float64(cows)

	return q, err
}
