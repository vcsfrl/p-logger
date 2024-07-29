package logger

import (
	"bytes"
	"time"
)

// Mocks

func mockNow() func() time.Time {
	return func() time.Time {
		return time.Date(2024, 5, 1, 3, 12, 3, 0, time.UTC)
	}
}

type MockWriter struct {
	bytes.Buffer
}

func (m *MockWriter) Close() error {
	return nil
}
