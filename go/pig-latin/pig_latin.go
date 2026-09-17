package piglatin

import (
	"regexp"
	"strings"
)

var rule1, _ = regexp.Compile("^(a|e|i|o|u|xr|yt).+")
var rule2, _ = regexp.Compile("^([^aeiou]+)(.+)")
var rule3, _ = regexp.Compile("^([^aeiou]*qu)(.+)")
var rule4, _ = regexp.Compile("^([^aeiou]+)(y)(.+)")

func Sentence(sentence string) string {
	var translated []string

	for word := range strings.FieldsSeq(sentence) {
		translated = append(translated, Word(word))
	}

	return strings.Join(translated, " ")
}

func Word(word string) string {
	if rule1.MatchString(word) {
		return word + "ay"
	} else if matches := rule4.FindStringSubmatch(word); matches != nil {
		return matches[2] + matches[3] + matches[1] + "ay"
	} else if matches := rule3.FindStringSubmatch(word); matches != nil {
		return matches[2] + matches[1] + "ay"
	} else if matches := rule2.FindStringSubmatch(word); matches != nil {
		return matches[2] + matches[1] + "ay"
	}

	return word
}
