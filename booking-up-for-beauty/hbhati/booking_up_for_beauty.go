package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// Input: 7/25/2019 13:45:00
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// Input: July 25, 2019 13:45:00
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// Input: Thursday, July 25, 2019 13:45:00
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
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
	anniversaryDate := fmt.Sprintf("09/15/%4d 00:00:00", time.Now().Year())
	return Schedule(anniversaryDate)
}
