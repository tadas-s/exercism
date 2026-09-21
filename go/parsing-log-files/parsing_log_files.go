package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var validLineRe = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)].*`)
var splitLogLineRe = regexp.MustCompile(`<[~*=-]*>`)
var quotedPasswordRe = regexp.MustCompile(`(?i)"[^"]*(password)[^"]*"`)
var endOfLineRe = regexp.MustCompile(`end-of-line\d+`)
var userNameRe = regexp.MustCompile(`User\s+(\S+)`)

func IsValidLine(text string) bool {
	return validLineRe.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogLineRe.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0

	for _, line := range lines {
		found := quotedPasswordRe.FindAllString(line, -1)
		count += len(found)
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRe.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var tagged []string

	for _, line := range lines {
		if userMatch := userNameRe.FindStringSubmatch(line); userMatch != nil {
			tagged = append(tagged, fmt.Sprintf("[USR] %s %s", userMatch[1], line))
		} else {
			tagged = append(tagged, line)
		}
	}

	return tagged
}
