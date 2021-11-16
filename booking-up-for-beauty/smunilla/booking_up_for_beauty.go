package booking

import (
	"fmt"
	"time"
)

func parseDate(date string) time.Time {
	var layouts = []string{
		"1/2/2006 15:04:05",
		"January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t
		}
	}

	panic(fmt.Sprintf("Problem parsing date: %s", date))
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	return parseDate(date)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return time.Now().After(parseDate(date))
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	apptHour := parseDate(date).Hour()

	return apptHour >= 12 && apptHour <= 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	appt := parseDate(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		appt.Weekday(),
		appt.Month(),
		appt.Day(),
		appt.Year(),
		appt.Hour(),
		appt.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
