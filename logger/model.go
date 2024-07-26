package logger

import "time"

type Severity string

const (
	SeverityDebug   Severity = "DEBUG"
	SeverityInfo    Severity = "INFO"
	SeverityWarning Severity = "WARN"
	SeverityError   Severity = "ERROR"
	SeverityDefault Severity = SeverityInfo
)

var severities = []Severity{SeverityDefault, SeverityDebug, SeverityInfo, SeverityWarning, SeverityError}

// Message is a DTO that represents a log message containing the fields the user can set.
type Message struct {
	Content    string
	Attributes map[string]string
	Tags       []string
}

// logMessage is the struct that will be used to log messages.
// It is not exported because it should be used only internally.
type logMessage struct {
	Timestamp time.Time
	Severity  Severity
	Message
}
