package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch {
	case distance <= 1.0:
		return 10
	case distance <= 5.0:
		return 5
	case distance <= 10.0:
		return 1
	default:
		return 0
	}
}
