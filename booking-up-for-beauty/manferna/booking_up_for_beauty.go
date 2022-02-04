package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	// layout constant
	const layout = "1/02/2006 15:04:05"

	//dateFormat := strings.ReplaceAll(date, "/", "-")
	appointmentTime, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	//fmt.Println(date)
	//fmt.Println(appointmentTime)
	return appointmentTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	currentTime := time.Now()
	appointmentTime, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return appointmentTime.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {

	const layout = "Monday, January 2, 2006 15:04:05"
	const start = 12
	const end = 18

	appointmentTime, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	if (appointmentTime.Hour() >= start) && (appointmentTime.Hour() < end) {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {

	const layout = "1/2/2006 15:04:05"
	appointmentDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s", appointmentDate.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	// September 15th in 2012
	// Output: 2020-09-15 00:00:00 +0000 UTC

	anniversaryDateNow := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)

	return anniversaryDateNow
}
