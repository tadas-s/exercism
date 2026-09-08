package armstrongnumbers

import (
	"math"
)

func IsNumber(n int) bool {
	numOfDigits := int(math.Floor(math.Log10(float64(n)))) + 1
	sumOfPows := 0

	for i := range numOfDigits {
		digit := (n % int(math.Pow(10, float64(i+1)))) / int(math.Pow(10, float64(i)))
		sumOfPows += int(math.Pow(float64(digit), float64(numOfDigits)))
	}

	return sumOfPows == n
}
