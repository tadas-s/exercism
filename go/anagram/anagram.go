package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string

	sortedSubject := sortedRunes(subject)

	for _, candidate := range candidates {
		if strings.ToLower(subject) != strings.ToLower(candidate) && sortedSubject == sortedRunes(candidate) {
			result = append(result, candidate)
		}
	}

	return result
}

func sortedRunes(s string) string {
	runes := []rune(strings.ToLower(s))

	slices.SortFunc(runes, func(a, b rune) int {
		return strings.Compare(string(a), string(b))
	})

	return string(runes)
}
