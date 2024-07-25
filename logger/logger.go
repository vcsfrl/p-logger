package logger

import (
	"bytes"
	"fmt"
	"time"
)

type Logger struct {
	output *bytes.Buffer
	now    func() time.Time
}

func (l *Logger) Log(message string) {
	l.output.WriteString(fmt.Sprintf("%s : %s", l.now().Format(time.RFC3339), message))
}

func NewLogger(output *bytes.Buffer) *Logger {
	return &Logger{
		output: output,
		now:    time.Now,
	}
}
