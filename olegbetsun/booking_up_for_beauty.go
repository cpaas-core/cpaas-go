package booking

import (
	"time"
	"strconv"
	"fmt"
	)

func ParseDate(layout string, date string) time.Time{
	t,err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	return t
}


// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t:= ParseDate(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t := ParseDate(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t:= ParseDate(layout, date)
	noon := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, t.Nanosecond(), t.Location())
	eve := time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, t.Nanosecond(), t.Location())
	return t.After(noon) && t.Before(eve)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t:= ParseDate(layout, date)
	return "You have an appointment on "+t.Weekday().String()+", "+t.Month().String()+" "+strconv.Itoa(t.Day())+", "+strconv.Itoa(t.Year())+", at "+strconv.Itoa(t.Hour())+":"+strconv.Itoa(t.Minute())+"."

}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
