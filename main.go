package main

import (
	"fmt"

	"New-Year2/calendar"
)

func main() {
	var date string

	_, err := fmt.Scan(&date)
	if err != nil {
		fmt.Println(err)
		return
	}

	days, err := calendar.NewYear(date)
	if err != nil {
		fmt.Println("Ошибка: некорректный формат даты. Используйте ДД.ММ.ГГГГ")
		return
	}
	fmt.Printf("До New year %d дней", days)
}
