package objects

import "time"

type BusinessHours struct {
	Periods []*BusinessHoursPeriod `json:"periods"`
}

type BusinessHoursPeriod struct {
	DayOfWeek      string     `json:"day_of_week,omitempty"`
	StartLocalTime *time.Time `json:"start_local_time,omitempty"`
	EndLocalTime   *time.Time `json:"end_local_time,omitempty"`
}
