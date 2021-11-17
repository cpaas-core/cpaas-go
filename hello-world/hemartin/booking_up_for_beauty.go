package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return parsedTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()

	appointmentDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return now.After(appointmentDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	appointmentDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	if 12 <= appointmentDate.Hour() && appointmentDate.Hour() < 18 {
		return true
	}

	return false

}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	appointmentDate := Schedule(date)
	return appointmentDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
