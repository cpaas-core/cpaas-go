package booking

import (
	"fmt"
	"time"
)

func parseDate(layout string, dateStr string) time.Time {
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Printf("An error occured while parsing time: %s\n", err)
	}

	return parsedDate
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout string = "1/02/2006 15:04:05"
	return parseDate(layout, date)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout string = "January 2, 2006 15:04:05"
	apt := parseDate(layout, date)
	now := time.Now()
	return apt.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const layout string = "Monday, January 2, 2006 15:04:05"
	hour := parseDate(layout, date).Hour()
	return hour >= 12 && hour <= 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const parseLayout string = "1/2/2006 15:04:05"
	const formatLayout string = "Monday, January 2, 2006, at 15:04"
	apt := parseDate(parseLayout, date)
	return fmt.Sprintf("You have an appointment on %s.", apt.Format(formatLayout))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, 9, 15, 0, 0, 0, 0, time.UTC)

}
