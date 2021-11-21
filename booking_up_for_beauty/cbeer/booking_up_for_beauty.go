package booking

import (
	"time"
 	"fmt"
	)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// Input example: 7/25/2019 13:45:00
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// Input example: July 25, 2019 13:45:00
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// Input example: Thursday, July 25, 2019 13:45:00
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// Input example: 7/25/2019 13:45:00
	// Output example: Thursday, July 25, 2019, at 13:45
	t := Schedule(date)
	appointment := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", appointment)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	anniversaryDate := fmt.Sprintf("09/15/%4d 00:00:00", time.Now().Year())
	return Schedule(anniversaryDate)
}
