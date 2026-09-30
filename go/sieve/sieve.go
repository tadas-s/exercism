package sieve

func Sieve(limit int) []int {
	composite := make([]bool, limit+1)

	number := 2

	for number*number <= limit {
		multiple := number * 2

		for multiple <= limit {
			composite[multiple] = true
			multiple += number
		}

		number++

		for number <= limit && composite[number] == true {
			number++
		}
	}

	// From benchmarks it looks like it's worth going through
	// the sieve and counting primes found to allocate
	// exact array to store results. Otherwise re-allocations
	// by `append` does add noticeable time and memory overhead.
	foundPrimes := 0

	for i := 2; i < len(composite); i++ {
		if !composite[i] {
			foundPrimes++
		}
	}

	var primes = make([]int, foundPrimes)
	n := 0

	for i := 2; i < len(composite); i++ {
		if !composite[i] {
			primes[n] = i
			n++
		}
	}

	return primes
}
