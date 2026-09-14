package swiftscheduling

import (
	"strconv"
	"strings"
	"time"
)

const dateTimeLayout = "2006-01-02T15:04:05"

func DeliveryDate(start, delivery string) string {
	startTime, err := time.Parse(dateTimeLayout, start)
	var quarter, month int

	if err != nil {
		panic(err)
	}

	if quarterString, ok := strings.CutPrefix(delivery, "Q"); ok {
		quarter, _ = strconv.Atoi(quarterString)
		delivery = "Q"
	}

	if monthString, ok := strings.CutSuffix(delivery, "M"); ok {
		month, _ = strconv.Atoi(monthString)
		delivery = "M"
	}

	switch delivery {
	case "NOW":
		return deliverNOW(startTime)
	case "ASAP":
		return deliverASAP(startTime)
	case "EOW":
		return deliverEOW(startTime)
	case "M":
		return deliverM(startTime, month)
	case "Q":
		return deliverQ(startTime, quarter)
	default:
		panic("Unsupported `delivery` argument")
	}
}

func deliverNOW(start time.Time) string {
	return (start.Add(2 * time.Hour)).Format(dateTimeLayout)
}

func deliverASAP(start time.Time) string {
	if start.Hour() < 12 {
		beforeEOD := time.Date(start.Year(), start.Month(), start.Day(), 17, 0, 0, 0, time.UTC)
		return beforeEOD.Format(dateTimeLayout)
	}

	beforeNextDayAfternoon := time.Date(start.Year(), start.Month(), start.Day()+1, 13, 0, 0, 0, time.UTC)
	return beforeNextDayAfternoon.Format(dateTimeLayout)
}

func deliverEOW(start time.Time) string {
	if start.Weekday() < time.Thursday {
		beforeFriday := time.Date(start.Year(), start.Month(), start.Day(), 17, 0, 0, 0, time.UTC)
		beforeFriday = beforeFriday.Add(time.Hour * 24 * time.Duration(time.Thursday-start.Weekday()+1))
		return beforeFriday.Format(dateTimeLayout)
	}

	beforeSunday := time.Date(start.Year(), start.Month(), start.Day(), 20, 0, 0, 0, time.UTC)
	beforeSunday = beforeSunday.Add(time.Hour * 24 * time.Duration(time.Saturday-start.Weekday()+1))
	return beforeSunday.Format(dateTimeLayout)
}

func deliverM(start time.Time, month int) string {
	beforeMonth := time.Date(start.Year(), time.Month(month), 1, 8, 0, 0, 0, time.UTC)

	if start.Month() >= time.Month(month) {
		beforeMonth = beforeMonth.AddDate(1, 0, 0)
	}

	switch beforeMonth.Weekday() {
	case time.Sunday:
		beforeMonth = beforeMonth.AddDate(0, 0, 1)
	case time.Saturday:
		beforeMonth = beforeMonth.AddDate(0, 0, 2)
	default:
		break
	}

	return beforeMonth.Format(dateTimeLayout)
}

func deliverQ(start time.Time, quarter int) string {
	thisQuarter := int((start.Month() - 1) / 3)

	beforeQuarter := time.Date(start.Year(), time.Month(quarter*3+1), 1, 8, 0, 0, 0, time.UTC)

	if thisQuarter > (quarter - 1) {
		beforeQuarter = beforeQuarter.AddDate(1, 0, 0)
	}

	beforeQuarter = beforeQuarter.AddDate(0, 0, -1)

	switch beforeQuarter.Weekday() {
	case time.Sunday:
		beforeQuarter = beforeQuarter.AddDate(0, 0, -2)
	case time.Saturday:
		beforeQuarter = beforeQuarter.AddDate(0, 0, -1)
	default:
		break
	}

	return beforeQuarter.Format(dateTimeLayout)
}
