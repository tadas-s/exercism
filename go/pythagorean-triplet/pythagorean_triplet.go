package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var triplets []Triplet

	for a := 1; a <= p/3; a++ {
		for b := a + 1; b <= p/2; b++ {
			c := p - a - b

			if a*a+b*b == c*c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}

	return triplets
}
