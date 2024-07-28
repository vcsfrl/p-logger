package logger

// LevelLogger is a logger that logs messages with a specific severity level.
// It is a wrapper around the Logger struct, ad adds some utility methods.
type LevelLogger struct {
	*Logger
}

func NewLevelLogger(logger *Logger) *LevelLogger {
	return &LevelLogger{Logger: logger}
}

func (l *LevelLogger) Info(message string) {
	l.Log(SeverityInfo, Message{Content: message})
}

func (l *LevelLogger) Debug(message string) {
	l.Log(SeverityDebug, Message{Content: message})
}

func (l *LevelLogger) Warn(message string) {
	l.Log(SeverityWarning, Message{Content: message})
}

func (l *LevelLogger) Error(message string) {
	l.Log(SeverityError, Message{Content: message})
}

func (l *LevelLogger) Transaction(id string, attributes map[string]string) {
	message := Message{Content: "Transaction", Attributes: attributes}
	message.Attributes["transaction_id"] = id
	l.Log(SeverityDefault, message)
}

func (l *LevelLogger) Close() error {
	return l.outputWriter.Close()
}
