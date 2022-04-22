package thefarm
import (
    "fmt"
    "errors"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
    count int
}
func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows",e.count )
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    if cows == 0 {
        return 0.0, errors.New("division by zero")
    }
	if cows < 0 {
        return 0.0, &SillyNephewError{cows}
        }
	
	fodder, err := weightFodder.FodderAmount()
        if err != nil {
        	if err == ErrScaleMalfunction {
        		fodder = fodder * 2
    }  else{
			return 0.0, err
                
    }
        }
		if fodder < 0 {
    		return 0.0, errors.New("negative fodder")
}
		return fodder / float64(cows),nil
}