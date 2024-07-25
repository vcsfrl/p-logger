package logger

import "bytes"

type Logger struct {
	output *bytes.Buffer
}

func (l *Logger) Log(message string) {
	l.output.WriteString(message)
}

func NewLogger(output *bytes.Buffer) *Logger {
	return &Logger{
		output: output,
	}
}
