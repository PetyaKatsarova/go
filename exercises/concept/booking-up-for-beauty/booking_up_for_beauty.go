// package booking

package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05" // layout is a hardcoded date string in the time library for a reference
	result, _ := time.Parse(layout, date)
	return result
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	result, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println("error parsing time:", err)
	}
	return result.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	d, err := time.Parse("Monday, January 2, 2006 15:04:05", date) //Thursday, May 13, 2010 20:32:00
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	return d.Hour() >= 12 && d.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	result, err := time.Parse("1/2/2006 15:04:05", date) //6/6/2005 10:30:00
	if err != nil {
		fmt.Println("error parsing time:", err)
	}
	return fmt.Sprintf("You have an appointment on %s, %s, at %s.", result.Format("Monday"), result.Format("January 2, 2006"), result.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	fmt.Println(Description("7/6/2005 10:30:00"))
	fmt.Println(time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC))
	fmt.Println(AnniversaryDate())
	// fmt.Println(IsAfternoonAppointment("Thursday, May 13, 2010 20:32:00")) // false
	// fmt.Println(Schedule("2/2/2222 22:22:22"))
	// fmt.Println(time.Now())
	// fmt.Println("definitely not passed: false: ", HasPassed("December 9, 2112 11:59:59"))
}
