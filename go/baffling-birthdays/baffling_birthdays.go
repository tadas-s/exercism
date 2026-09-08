package bafflingbirthdays

import (
	"math/rand"
	"slices"
	"time"
)

func BirthdayCompare(a, b time.Time) int {
	if b.Month()-a.Month() != 0 {
		return int(b.Month() - a.Month())
	}

	return b.Day() - a.Day()
}

func SharedBirthday(dates []time.Time) bool {
	slices.SortFunc(dates, func(a, b time.Time) int {
		return BirthdayCompare(a, b)
	})

	for i := 0; i < len(dates)-1; i++ {
		if BirthdayCompare(dates[i], dates[i+1]) == 0 {
			return true
		}
	}

	return false
}

func RandomBirthdates(size int) []time.Time {
	birthdates := make([]time.Time, size)
	year := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := range size {
		birthdates[i] = time.Unix(rand.Int63n(endOfYear.Unix()-year.Unix())+year.Unix(), 0)
	}

	return birthdates
}

func EstimatedProbability(size int) float64 {
	experiments := 5000
	sharedBirthdays := 0

	for range experiments {
		birthdates := RandomBirthdates(size)

		if SharedBirthday(birthdates) {
			sharedBirthdays++
		}
	}

	return float64(sharedBirthdays) / float64(experiments) * 100.0
}
