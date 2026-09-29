package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleNewYear(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleNewYear))
	defer server.Close()

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		verifyLogic    func(t *testing.T, resp Response)
	}{
		{
			name:           "Успешный запрос",
			queryParams:    "?date=30.12.2026",
			expectedStatus: http.StatusOK,
			verifyLogic: func(t *testing.T, resp Response) {
				if resp.DaysLeft != 2 {
					t.Errorf("Ожидали 2 дня до Нового года для даты 30.12.2026, получили: %d", resp.DaysLeft)
				}
				if resp.Target != "30.12.2026" {
					t.Errorf("Ожидали target_date '30.12.2026', получили: %s", resp.Target)
				}
			},
		},
		{
			name:           "Запрос без указания даты (текущая дата)",
			queryParams:    "",
			expectedStatus: http.StatusOK,
			verifyLogic: func(t *testing.T, resp Response) {
				expectedToday := time.Now().Format("02.01.2006")
				if resp.Target != expectedToday {
					t.Errorf("Ожидали текущую дату '%s', получили: %s", expectedToday, resp.Target)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetURL := server.URL + tt.queryParams
			res, err := http.Get(targetURL)
			if err != nil {
				t.Fatalf("Не удалось выполнить HTTP-запрос: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Ожидали статус-код %d, получили %d", tt.expectedStatus, res.StatusCode)
			}

			var responseBody Response
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			if err != nil {
				t.Fatalf("Не удалось распарсить JSON-ответ: %v", err)
			}

			tt.verifyLogic(t, responseBody)
		})
	}
}
