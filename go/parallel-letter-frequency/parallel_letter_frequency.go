package parallelletterfrequency

import "unicode"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Merge incorporates frequency counts from the <other> map
func (fmap *FreqMap) Merge(other FreqMap) {
	for r, otherCount := range other {
		(*fmap)[r] += otherCount
	}
}

func (fmap *FreqMap) Count(text string) {
	for _, r := range text {
		if !unicode.IsLetter(r) {
			continue
		}
		(*fmap)[unicode.ToLower(r)] += 1
	}
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	counts := FreqMap{}
	counts.Count(text)
	return counts
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	totals := FreqMap{}
	workers := min(5, len(texts))

	inCh := make(chan string, len(texts))
	for _, text := range texts {
		inCh <- text
	}

	outCh := make(chan FreqMap, workers)

	for range workers {
		go func() {
			counts := FreqMap{}

			for job := range inCh {
				counts.Count(job)
			}

			outCh <- counts
		}()
	}

	close(inCh)

	for range workers {
		totals.Merge(<-outCh)
	}

	return totals
}
