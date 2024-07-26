package logger

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// OutputWriter defines an interface for writing logs.
// The write method should not return any errors and should be safe for concurrent use.
type OutputWriter interface {
	Write(message LogMessage)
}

// TextOutputWriter is an OutputWriter that writes logs in text format to an io.Writer.
type TextOutputWriter struct {
	writer io.Writer
}

func (t *TextOutputWriter) Write(logMessage LogMessage) {
	// We do not handle errors. The writers should be in a state that no errors occur.
	// This is how it is implemented in the standard library logger as well.
	// We could panic on error, but that would be a bit too harsh.
	// We could also log the error, with the standard library logger, but that would be
	// a logger in the logger (having the same flaw).
	_, _ = t.writer.Write([]byte(
		fmt.Sprintf(
			"%s :: %s :: %s :: %s :: %v \n",
			logMessage.Timestamp.Format(time.RFC3339),
			logMessage.Severity,
			logMessage.Content,
			strings.TrimLeft(fmt.Sprint(logMessage.Attributes), "map"),
			logMessage.Tags,
		),
	))
}

// MultiOutputWriter is an OutputWriter that writes logs to multiple writers concurrently.
type MultiOutputWriter struct {
	writers []OutputWriter
}

// Write sends the log message to all writers concurrently.
func (m *MultiOutputWriter) Write(message LogMessage) {
	var wg sync.WaitGroup
	for _, writer := range m.writers {
		wg.Add(1)
		go func() {
			writer.Write(message)
			wg.Done()
		}()
	}
	wg.Wait()
}

// NewMultiOutputWriter constructor for MultiOutputWriter.
func NewMultiOutputWriter(writers ...OutputWriter) *MultiOutputWriter {
	return &MultiOutputWriter{writers: writers}
}
