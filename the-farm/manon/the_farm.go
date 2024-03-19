package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, n int) (float64, error) {
	if amount, err := fodderCalculator.FodderAmount(n); err != nil {
		return 0, err
	} else if factor, err := fodderCalculator.FatteningFactor(); err != nil {
		return 0, err
	} else {
		return amount / float64(n) * factor, nil
	}
}

func ValidateInputAndDivideFood(calculator FodderCalculator, n int) (float64, error) {
	if n > 0 {
		return DivideFood(calculator, n)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.message)
}

func ValidateNumberOfCows(n int) error {
	switch {
	case n < 0:
		return &InvalidCowsError{n, "there are no negative cows"}
	case n == 0:
		return &InvalidCowsError{n, "no cows don't need food"}
	default:
		return nil
	}
}
