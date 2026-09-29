package acronym

import (
	"strings"
	"unicode"
)

type scanState int

const (
	stateSeparator scanState = iota
	stateWord
)

// Abbreviate returns abbreviation/acronym of a phrase
func Abbreviate(s string) string {
	var abbreviation strings.Builder
	state := stateSeparator

	for _, c := range s {
		switch state {
		case stateSeparator:
			if unicode.IsLetter(c) {
				abbreviation.WriteRune(unicode.ToUpper(c))
				state = stateWord
			}
		case stateWord:
			if !unicode.IsLetter(c) && c != '\'' {
				state = stateSeparator
			}
		}
	}

	return abbreviation.String()
}
