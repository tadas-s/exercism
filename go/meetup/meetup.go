package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	nthCounter := 0
	lastWeekday := -1

	for currentDay := 1; currentDay <= 31; currentDay++ {
		date := time.Date(year, month, currentDay, 12, 0, 0, 0, time.UTC)

		// If month was bumped - we're looping month with less than 31 days
		if date.Month() != month {
			break
		}

		if date.Weekday() != wDay {
			continue
		}

		switch {
		case wSched == Teenth && date.Day() >= 13 && date.Day() <= 19:
			return currentDay
		case wSched >= First && wSched <= Fourth:
			if nthCounter == int(wSched) {
				return currentDay
			}
			nthCounter++
		case wSched == Last:
			lastWeekday = currentDay
		}
	}

	return lastWeekday // for when wSched == Last
}
