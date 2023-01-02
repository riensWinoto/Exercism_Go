package main

import (
	"errors"
	"fmt"
)

var (
	ErrScaleMalfunction = errors.New("sensor error")
	NegativeFodder      = errors.New("negative fodder")
	DivisionZero        = errors.New("division by zero")
)

type WeightFodder interface {
	FodderAmount() (float64, error)
}

type SillyNephewError struct {
	cows int
}

func (silly *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", silly.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			if fodder < 0 {
				return 0, NegativeFodder
			}
			return (2 * fodder) / float64(cows), nil
		}
		return 0, err
	}

	if fodder < 0 {
		return 0, NegativeFodder
	}

	if cows == 0 {
		return 0, DivisionZero
	}

	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	return fodder / float64(cows), nil
}
