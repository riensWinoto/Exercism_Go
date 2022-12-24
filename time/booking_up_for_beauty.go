package main

import (
	"fmt"
	"time"
)

var (
	schedule            string = "7/25/2019 13:45:00"
	passedSchedule      string = "July 25, 2019 13:45:00"
	appointmentSchedule string = "Thursday, July 25, 2019 13:45:00"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	customLayout := "1/02/2006 15:04:05"
	fixedSchedule, err := time.Parse(customLayout, date)
	if err != nil {
		fmt.Println(err)
	}
	return fixedSchedule
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	customLayout := "January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(customLayout, date)
	if err != nil {
		fmt.Println(err)
	}
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	customLayout := "Monday, January 2, 2006 15:04:05"
	parsedAppointmentDate, err := time.Parse(customLayout, date)
	if err != nil {
		fmt.Println(err)
	}
	return parsedAppointmentDate.Hour() >= 12 && parsedAppointmentDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	customLayout := "1/2/2006 15:04:05"
	parsedDate, err := time.Parse(customLayout, date)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.",
		parsedDate.Weekday(), parsedDate.Month(), parsedDate.Day(), parsedDate.Year(), parsedDate.Hour(), parsedDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	fmt.Println(Schedule(schedule))
	fmt.Println(HasPassed(passedSchedule))
	fmt.Println(IsAfternoonAppointment(appointmentSchedule))
	fmt.Println(Description(schedule))
	fmt.Println(AnniversaryDate())
}
