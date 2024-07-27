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

// LogMessage is the struct that will be used to log messages.
// This is exported because users can define OutputWriters outside this library.
type LogMessage struct {
	Timestamp time.Time
	Severity  Severity
	Message
}

type ConfigWriter struct {
	Name   string         `json:"name"`
	Params map[string]any `json:"params"`
}

// Config is a configuration structure for building the logger.
type Config struct {
	Writers []ConfigWriter `json:"writers"`
}
