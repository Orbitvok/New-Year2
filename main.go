package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"New-Year2/calendar"
)

// Response определяет структуру машиночитаемого ответа сервера.
type Response struct {
	DaysLeft int    `json:"days_left"`
	Target   string `json:"target_date"`
}

// ErrorResponse определяет структуру ответа в случае ошибки.
type ErrorResponse struct {
	Error string `json:"error"`
}

// handleNewYear обрабатывает HTTP-запросы к API.
func handleNewYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Метод не поддерживается"})
		return
	}

	dateStr := r.URL.Query().Get("date")

	if dateStr == "" {
		dateStr = time.Now().Format("02.01.2006")
	}

	days, err := calendar.NewYear(dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Некорректный формат даты. Используйте ДД.ММ.ГГГГ"})
		return
	}

	res := Response{
		DaysLeft: days,
		Target:   dateStr,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/api/newyear", handleNewYear)

	log.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
