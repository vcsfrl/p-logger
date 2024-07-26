package logger

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Logger struct {
	output *bytes.Buffer
	now    func() time.Time
}

type Severity string

const (
	SeverityDebug   Severity = "DEBUG"
	SeverityInfo    Severity = "INFO"
	SeverityWarning Severity = "WARN"
	SeverityError   Severity = "ERROR"
	SeverityDefault Severity = SeverityInfo
)

var severities = []Severity{SeverityDefault, SeverityDebug, SeverityInfo, SeverityWarning, SeverityError}

type Message struct {
	Content    string
	Attributes map[string]string
	Tags       []string
}

type logMessage struct {
	Timestamp time.Time
	Severity  Severity
	Message
}

// Log builds a logMessage and sends it to the output writer..
// If the severity is not valid, it will default to SeverityDefault.
// This method should not return any errors.
func (l *Logger) Log(severity Severity, message Message) {
	if !slices.Contains(severities, severity) {
		severity = SeverityDefault
	}

	// Build the log message. Timestamp is set by the logger.
	logeMessage := logMessage{
		Timestamp: l.now(),
		Severity:  severity,
		Message:   message,
	}

	l.output.WriteString(
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
		output: output,
		now:    time.Now,
	}
}
