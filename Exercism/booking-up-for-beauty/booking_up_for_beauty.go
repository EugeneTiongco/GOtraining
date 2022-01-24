package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	d, _ := time.Parse("1/02/2006 15:04:05", date)
	return d
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	d, _ := time.Parse("January 2, 2006 15:04:05", date)

	return d.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	d, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if d.Hour() >= 12 && d.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	d, _ := time.Parse("1/2/2006 15:04:05", date)
	msg := "You have an appointment on " + d.Format("Monday, January 2, 2006, at 15:04.")
	return msg
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}
