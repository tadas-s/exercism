package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0

	for n := 1; n < limit; n++ {
		for _, divisor := range divisors {
			if divisor > 0 && n%divisor == 0 {
				sum += n
				break
			}
		}
	}

	return sum
}
