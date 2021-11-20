package booking

import "time"


// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	myDate, _ := time.Parse("1/2/2006 15:04:05", date)
	return myDate
	
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	today := time.Now()
	//myDateString := "July 25, 2019 13:45:00"
	myDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	return today.After(myDate)
	
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	myDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return myDate.Hour() >= 12 && myDate.Hour() < 18
	
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	myDate := Schedule(date)
	return myDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, 9, 15, 0, 0, 0, 0, time.UTC)
	
}


