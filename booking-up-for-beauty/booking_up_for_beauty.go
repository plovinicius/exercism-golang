package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// panic("Please implement the Schedule function")
	formatToParse := "1/_2/2006 15:04:05"
	formattedTime,_ := time.Parse(formatToParse, date)

	return formattedTime

}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// panic("Please implement the HasPassed function")
	current := time.Now()
	toCheck,_ := time.Parse("January _2, 2006 15:04:05", date)

	return current.After(toCheck)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// panic("Please implement the IsAfternoonAppointment function")
	toCheck,_ := time.Parse("Monday, January _2, 2006 15:04:05", date)

	if toCheck.Hour() >= 12 && toCheck.Hour() <= 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// panic("Please implement the Description function")

	toCheck,_ := time.Parse("1/_2/2006 15:04:05", date)
	formattedDate := toCheck.Format("Monday, January 2, 2006")
	formattedTime := toCheck.Format("15:04")

	return fmt.Sprintf("You have an appointment on %v, at %v.", formattedDate, formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	// panic("Please implement the AnniversaryDate function")
	currentYear := time.Now().Year()
	datetime,_ := time.Parse("_2/01/2006 15:04:05", fmt.Sprintf("15/09/%v 00:00:00", currentYear))

	return datetime
}
