package pangram

import "strings"

func IsPangram(input string) bool {
	lowercased := strings.ToLower(input)

	for c := 'a'; c <= 'z'; c++ {
		if !strings.ContainsRune(lowercased, c) {
			return false
		}
	}

	return true
}
