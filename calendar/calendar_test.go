package calendar

import "testing"

func TestNewYear(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		wantErr  bool
	}{
		{
			name:     "Начало календарного дня",
			input:    "01.01.2026",
			expected: 365,
			wantErr:  false,
		},
		{
			name:     "Конец календарного дня",
			input:    "31.12.2026",
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "Високосный год",
			input:    "28.02.2024",
			expected: 308,
			wantErr:  false,
		},
		{
			name:     "Обычный год",
			input:    "28.02.2026",
			expected: 307,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewYear(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewYear(%q) вернул ошибку: %v", tt.input, err)
				return
			}
			if actual != tt.expected {
				t.Errorf("NewYear(%q) = %d; ожидали %d", tt.input, actual, tt.expected)
			}
		})
	}
}
