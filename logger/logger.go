package logger

import (
	"os"
	"slices"
	"time"
)

type Logger struct {
	outputWriter OutputWriter
	now          func() time.Time
}

// Log builds a LogMessage and sends it to the outputWriter writer.
// If the severity is not valid, it will default to SeverityDefault.
// This method should not return any errors.
func (l *Logger) Log(severity Severity, message Message) {
	if !slices.Contains(severities, severity) {
		severity = SeverityDefault
	}

	// Build the log message. Timestamp is set by the logger.
	logMessage := LogMessage{
		Timestamp: l.now(),
		Severity:  severity,
		Message:   message,
	}

	l.outputWriter.Write(logMessage)
}

func NewLogger(output OutputWriter) *Logger {
	if output == nil {
		output = &TextOutputWriter{writer: os.Stdout}
	}

	return &Logger{
		outputWriter: output,
		now:          time.Now,
	}
}
