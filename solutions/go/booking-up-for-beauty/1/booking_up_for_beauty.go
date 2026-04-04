package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedDate, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic("error parsing date")
	}

	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic("error parsing date")
	}
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic("error parsing date")
	}

	if parsedDate.Hour() >= 12 && parsedDate.Hour() < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic("erro parsing date")
	}

	return fmt.Sprintf(
		"You have an appointment on %s",
		parsedDate.Format("Monday, January 2, 2006, at 15:04."),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	parsedTime, err := time.Parse(
		"January 2 2006",
		fmt.Sprintf("September 15 %d", time.Now().Year()),
	)
	if err != nil {
		panic("error parsing date")
	}

	return parsedTime
}
