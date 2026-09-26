package calendar

import (
	"time"
)

func NewYear(date string) (int, error) {

	parse := "02.01.2006"

	dateParse, err := time.Parse(parse, date)
	if err != nil {
		return 0, err
	}

	nextYear := dateParse.Year() + 1

	newYearTime := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, dateParse.Location())
	days := int(newYearTime.Sub(dateParse) / (24 * time.Hour))

	return days, nil
}
