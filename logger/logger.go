package logger

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Logger struct {
	outputWriter *bytes.Buffer
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
	logeMessage := LogMessage{
		Timestamp: l.now(),
		Severity:  severity,
		Message:   message,
	}

	l.outputWriter.WriteString(
		fmt.Sprintf(
			"%s :: %s :: %s :: %s :: %v \n",
			logeMessage.Timestamp.Format(time.RFC3339),
			logeMessage.Severity,
			logeMessage.Content,
			strings.TrimLeft(fmt.Sprint(logeMessage.Attributes), "map"),
			logeMessage.Tags,
		))
}

func NewLogger(output *bytes.Buffer) *Logger {
	return &Logger{
		outputWriter: output,
		now:          time.Now,
	}
}
