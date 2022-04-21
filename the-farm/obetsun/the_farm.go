package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	var amount, err = weightFodder.FodderAmount()

	switch {
	case cows == 0:
		return 0, errors.New("Division by zero")
	case cows < 0:
		return 0, &SillyNephewError{cows}
	case err != nil:
		{
			if err == ErrScaleMalfunction {
				if amount < 0 {
					return 0, errors.New("Negative fodder")
				}
				var doubleAmount = amount * 2
				return doubleAmount / float64(cows), nil
			} else {
				return 0, err
			}
		}
	default:
		{
			if amount < 0 {
				return 0, errors.New("Negative fodder")
			}
			return amount / float64(cows), nil
		}
	}

}
