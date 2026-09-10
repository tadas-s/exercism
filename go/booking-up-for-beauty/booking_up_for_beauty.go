package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsed, err := time.Parse("1/2/2006 15:4:5", date)

	if err == nil {
		return parsed
	}

	panic(err)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsed, err := time.Parse("January 2, 2006 15:4:5", date)

	if err == nil {
		return time.Now().After(parsed)
	}

	panic(err)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsed, err := time.Parse("Monday, January 2, 2006 15:4:5", date)

	if err == nil {
		return parsed.Hour() >= 12 && parsed.Hour() < 18
	}

	panic(err)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsed, err := time.Parse("1/2/2006 15:4:5", date)

	if err == nil {
		return parsed.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	}

	panic(err)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
