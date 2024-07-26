package logger

import "time"

// Mocks

func mockNow() func() time.Time {
	return func() time.Time {
		return time.Date(2024, 5, 1, 3, 12, 3, 0, time.UTC)
	}
}
