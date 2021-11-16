package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	myDate, _ := time.Parse("1/2/2006 15:04:05", date)
	return myDate
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	today := time.Now()
	//myDateString := "July 25, 2019 13:45:00"
	myDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	return today.After(myDate)
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	myDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if myDate.Hour() >= 12 && myDate.Hour() < 18 {
		return true
	}
	if myDate.Hour() == 18 && myDate.Minute() == 00 {
		return true
	}
	return false
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	myDate := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", myDate.Weekday(), myDate.Month(), myDate.Day(), myDate.Year(), myDate.Hour(), myDate.Minute())
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	myDate, _ := time.Parse("2006-01-02", "2021-09-15")
	return myDate
	panic("Please implement the AnniversaryDate function")
}


