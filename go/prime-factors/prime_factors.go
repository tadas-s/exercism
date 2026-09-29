package primefactors

func Factors(n int64) []int64 {
	var factors []int64

	for factor := int64(2); factor <= n; factor++ {
		for n%factor == 0 {
			factors = append(factors, factor)
			n = n / factor
		}
	}

	return factors
}
