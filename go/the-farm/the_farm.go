package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
	fa, err := calc.FodderAmount(numberOfCows)

	if err != nil {
		return 0.0, err
	}

	ff, err := calc.FatteningFactor()

	if err != nil {
		return 0.0, err
	}

	return fa / float64(numberOfCows) * ff, nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows < 1 {
		return 0.0, errors.New("invalid number of cows")
	}

	return DivideFood(calc, numberOfCows)
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (ice InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.numberOfCows, ice.message)
}

func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalidCowsError{numberOfCows, "there are no negative cows"}
	case numberOfCows == 0:
		return &InvalidCowsError{numberOfCows, "no cows don't need food"}
	default:
		return nil
	}
}
