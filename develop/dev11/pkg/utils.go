package pkg

import (
	"time"
)

func ValidateDate(date string) (time.Time, bool) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, false
	}
	return d, true
}

func GetEventsForWeek(startDate time.Time) time.Time {
	endDate := startDate.AddDate(0, 0, 7)
	return endDate
}

func GetEventsForMonth(startDate time.Time) time.Time {
	endDate := startDate.AddDate(0, 1, 0)
	return endDate
}
