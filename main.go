package main

import (
	"fmt"
	"log"

	"New-Year2/calendar"
)

func main() {
	var dateStr string

	_, err := fmt.Scan(&dateStr)
	if err != nil {
		log.Fatalf("Ошибка ввода данных: %v", err)
	}

	days, err := calendar.NewYear(dateStr)
	if err != nil {
		log.Fatalf("Ошибка обработки даты: %v", err)
	}

	fmt.Printf("До Нового года осталось дней: %d\n", days)
}
