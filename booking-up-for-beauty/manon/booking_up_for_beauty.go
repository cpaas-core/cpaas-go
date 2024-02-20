package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	schedule, _ := time.Parse(layout, date)
	return schedule
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return time.Now().After(parsedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	hour := parsedDate.Hour()
	return (12 <= hour) && (hour < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "You have an appointment on Monday, January 2, 2006, at 15:04."
	return Schedule(date).Format(layout)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
