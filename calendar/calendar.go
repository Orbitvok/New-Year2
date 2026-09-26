package calendar

import (
	"time"
)

// NewYear вычисляет количество полных дней до наступления следующего Нового года.
func NewYear(date string) (int, error) {

	parse := "02.01.2006"

	parsedDate, err := time.Parse(parse, date)
	if err != nil {
		return 0, err
	}

	nextYear := parsedDate.Year() + 1

	newYearTime := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, parsedDate.Location())
	days := int(newYearTime.Sub(parsedDate) / (24 * time.Hour))

	return days, nil
}
