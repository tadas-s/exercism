package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	frequencies := Frequency{}

	split := strings.FieldsFunc(phrase, func(c rune) bool {
		return (unicode.IsSpace(c) || unicode.IsPunct(c)) && !(c == '\'')
	})

	for _, possiblyWord := range split {
		possiblyWord = strings.Trim(strings.ToLower(possiblyWord), "'")

		if len(possiblyWord) == 1 && !unicode.IsLetter(rune(possiblyWord[0])) && !unicode.IsDigit(rune(possiblyWord[0])) {
			continue
		}

		if len(possiblyWord) > 0 {
			frequencies[possiblyWord]++
		}
	}

	return frequencies
}
