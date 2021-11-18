package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const longForm = "1/2/2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		panic("Unable to parse date")
	}
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const longForm = "January 2, 2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		panic("Unable to parse date")
	}
	if t.Before(time.Now()) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const longForm = "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		panic("Unable to parse date")
	}

	if 12 <= t.Hour() && t.Hour() < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const longForm = "1/2/2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		panic("Unable to parse date")
	}

	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
