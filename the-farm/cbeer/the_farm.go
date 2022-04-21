package thefarm

import (
	"errors"
	"fmt"
	)

// SillyNephewError returned when the number of cows is negative.
type SillyNephewError struct {
	cowCount int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cowCount)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows < 0 {
		return 0.0, &SillyNephewError{cowCount: cows}
	} else if cows == 0 {
		return 0.0, errors.New("Division by zero")
	}
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			fodder *= 2.0
		} else {
			return 0.0, err
		}
	}
	if fodder < 0 {
		return 0.0, errors.New("Negative fodder")
	}
	return fodder / float64(cows), nil
}
