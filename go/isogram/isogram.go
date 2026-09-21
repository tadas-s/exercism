package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	rs := []rune(strings.ToLower(word))

	for i, r := range rs {
		if !unicode.IsLetter(r) {
			continue
		}

		for j := i + 1; j < len(rs); j++ {
			if r == rs[j] {
				return false
			}
		}
	}

	return true
}
