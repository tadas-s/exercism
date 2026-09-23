package nthprime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be > 0")
	}

	sequence := 1
	number := 2

	for sequence < n {
		for notPrime := true; notPrime; notPrime = !isPrime(number) {
			number++
		}

		sequence++
	}

	return number, nil
}

func isPrime(n int) bool {
	switch {
	case n < 1:
		return false
	case n == 1:
		return false
	case n == 2 || n == 3:
		return true
	case n%2 == 0 || n%3 == 0:
		return false
	default:
		for divisor := 5; divisor*divisor <= n; divisor += 6 {
			if n%divisor == 0 || n%(divisor+2) == 0 {
				return false
			}
		}

		return true
	}
}
