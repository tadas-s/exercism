package grains

import "fmt"

func Square(number int) (uint64, error) {
	if 1 > number || number > 64 {
		return 0, fmt.Errorf("square number %d out of range 1..64", number)
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
