package booking

import (
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// input format "7/25/2019 13:45:00"
	const layout = "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// input format "July 25, 2019 13:45:00"
	const layout = "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// input format "Thursday, July 25, 2019 13:45:00"
	const layout = "Monday, January 1, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return 12 <= t.Hour() && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// input format "7/25/2019 13:45:00"
	const inputLayout = "1/2/2006 15:04:05"
	const outputLayout = "Monday, January 2, 2006, at 15:04."
	t, _ := time.Parse(inputLayout, date)
	return "You have an appointment on " + t.Format(outputLayout)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	const layout = "2006-01-02"
	s := strconv.Itoa(time.Now().Year()) + "-09-15"
	t, _ := time.Parse(layout, s)
	return t
}
