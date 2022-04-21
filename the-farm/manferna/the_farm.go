package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError
type SillyNephewError struct {
	errorMessage string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %s cows", e.errorMessage)
}

func SillyNephewNegativeCows(cows int) error {
	return &SillyNephewError{
		errorMessage: fmt.Sprintf("%d", cows),
	}
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	// Before call fodderAmount, check for the cows!
	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0, SillyNephewNegativeCows(cows)
	}

	// We have cows!
	fodderAmount, err := weightFodder.FodderAmount()
	// Manage error
	if err != nil {
		if err == ErrScaleMalfunction {
			if fodderAmount < 0 {
				return 0.0, errors.New("negative fodder")
			} else {
				return (fodderAmount * 2.0) / float64(cows), nil
			}
		} else {
			return 0.0, err
		}
	}
	// Manage fodderAmount negative
	if fodderAmount < 0 {
		return 0.0, errors.New("negative fodder")
	}
	// Happy path
	return (fodderAmount / float64(cows)), nil
}
