package palindromeproducts

import (
	"errors"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	minP := Product{0, nil}
	maxP := Product{0, nil}

	if fmin > fmax {
		return minP, maxP, errors.New("min must be <= max")
	}

	for a := fmin; a <= fmax; a++ {
		for b := a; b <= fmax; b++ {
			p := a * b

			if isPalindrome(p) {
				if minP.Product == 0 || p < minP.Product {
					minP = Product{p, [][2]int{{a, b}}}
				} else if p == minP.Product {
					minP.Factorizations = append(minP.Factorizations, [2]int{a, b})
				}

				if maxP.Product == 0 || p > maxP.Product {
					maxP = Product{p, [][2]int{{a, b}}}
				} else if p == maxP.Product {
					maxP.Factorizations = append(maxP.Factorizations, [2]int{a, b})
				}
			}
		}
	}

	return minP, maxP, nil
}

func isPalindrome(input int) bool {
	original := input
	negative := input < 0
	var reverse int

	if negative {
		input *= -1
	}

	for input > 0 {
		reverse = reverse*10 + input%10
		input /= 10
	}

	if negative {
		reverse *= -1
	}

	return reverse == original
}
