package vetDomain

import "time"

type WorkDaySchedule struct {
	Day              time.Weekday
	WorkDayHourRange HourRange
	BreakHourRange   HourRange
}

type HourRange struct {
	StartHour int
	EndHour   int
}
