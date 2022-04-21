package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	numCows int
}

const sillyNephewErrorMsg = "silly nephew, there cannot be %d cows"

func (e SillyNephewError) Error() string {
	return fmt.Sprintf(sillyNephewErrorMsg, e.numCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0.0, SillyNephewError{cows}
	}

	fodderAmount, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction && fodderAmount > 0 {
		fodderAmount *= 2
	} else if err != nil && err != ErrScaleMalfunction {
		return 0.0, err
	}

	if fodderAmount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	return fodderAmount / float64(cows), nil
}
