package logger

import (
	"bytes"
)

type MemoryWriter struct {
	bytes.Buffer
}

func (m *MemoryWriter) Close() error {
	return nil
}
