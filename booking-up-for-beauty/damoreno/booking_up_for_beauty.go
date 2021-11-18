package booking

import (
	"fmt"
	"time"
)

func parseDate(layout, date string) time.Time {
	parsedDate, err := time.Parse(layout, date)

	if err != nil {
		panic("The suplied date is not parseable")
	}

	return parsedDate
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	return parseDate("1/02/2006 15:04:05", date)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return parseDate("January 2, 2006 15:04:05", date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	hour := parseDate("Monday, January 2, 2006 15:04:05", date).Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	parsedDate := parseDate("1/2/2006 15:04:05", date)

	return fmt.Sprintf("You have an appointment on %s.",
		parsedDate.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
