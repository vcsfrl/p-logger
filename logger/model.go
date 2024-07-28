package logger

import "time"

type Severity int

const (
	SeverityDebug Severity = iota + 1
	SeverityInfo
	SeverityWarning
	SeverityError
	SeverityDefault = SeverityInfo
)

var severityMap = map[string]Severity{
	"DEBUG":   SeverityDebug,
	"INFO":    SeverityInfo,
	"WARN":    SeverityWarning,
	"ERROR":   SeverityError,
	"DEFAULT": SeverityDefault,
}

func (s Severity) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[s-1]
}

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
	Writers     []ConfigWriter `json:"writers"`
	DefaultTags []string       `json:"default_tags"`
	MinSeverity string         `json:"min_severity"`
}
