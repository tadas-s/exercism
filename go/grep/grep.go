package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type MatchedLine struct {
	Number int
	Line   string
}

func Search(pattern string, flags, files []string) []string {
	numbered, namesOnly, caseInsensitive, invert, entireLineMatch := getFlags(flags)
	displayFileNames := len(files) > 1

	matches := map[string][]MatchedLine{}

	for _, file := range files {
		oneFileMatches := searchFile(pattern, file, caseInsensitive, invert, entireLineMatch)

		if len(oneFileMatches) > 0 {
			matches[file] = oneFileMatches
		}
	}

	return formatMatches(files, matches, displayFileNames, numbered, namesOnly)
}

func getFlags(flags []string) (numbered, namesOnly, caseInsensitive, invert, entireLineMatch bool) {
	for _, flag := range flags {
		switch flag {
		case "-n":
			numbered = true
		case "-l":
			namesOnly = true
		case "-i":
			caseInsensitive = true
		case "-v":
			invert = true
		case "-x":
			entireLineMatch = true
		default:
			break
		}
	}

	return
}

func searchFile(pattern string, fileName string, caseInsensitive, invert, entireLineMatch bool) []MatchedLine {
	var matches []MatchedLine
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for n := 1; scanner.Scan(); n++ {
		line := scanner.Text()

		if matchLine(pattern, line, caseInsensitive, invert, entireLineMatch) {
			matches = append(matches, MatchedLine{n, line})
		}
	}

	return matches
}

func matchLine(pattern, line string, caseInsensitive, invert, entireLineMatch bool) bool {
	result := false

	if caseInsensitive {
		pattern = strings.ToLower(pattern)
		line = strings.ToLower(line)
	}

	if entireLineMatch {
		result = pattern == line
	} else {
		result = strings.Contains(line, pattern)
	}

	if invert {
		return !result
	}

	return result
}

func formatMatches(files []string, matches map[string][]MatchedLine, displayFileNames, numbered, namesOnly bool) []string {
	var result []string

	for _, fileName := range files { // iterate over file name from args to preserve order in the results
		if fileMatches, found := matches[fileName]; found {
			if namesOnly {
				result = append(result, fileName)
				continue
			}

			for _, line := range fileMatches {
				if numbered && displayFileNames {
					result = append(result, fmt.Sprintf("%s:%d:%s", fileName, line.Number, line.Line))
				} else if numbered {
					result = append(result, fmt.Sprintf("%d:%s", line.Number, line.Line))
				} else if displayFileNames {
					result = append(result, fmt.Sprintf("%s:%s", fileName, line.Line))
				} else {
					result = append(result, line.Line)
				}
			}
		}
	}

	return result
}
