package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// OutputWriter defines an interface for writing logs.
// The write method should not return any errors and should be safe for concurrent use.
type OutputWriter interface {
	Write(message LogMessage)
	io.Closer
}

type OutputWriterConstructor func(attributes map[string]any) (OutputWriter, error)

// TEXT WRITER

// TextOutputWriter is an OutputWriter that writes logs in text format to an io.Writer.
type TextOutputWriter struct {
	writer io.WriteCloser
}

func (t *TextOutputWriter) Close() error {
	return t.writer.Close()
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

// NewTextStdoutWriter constructor for TextOutputWriter that writes to stdout.
func NewTextStdoutWriter(attributes map[string]any) (OutputWriter, error) {
	return &TextOutputWriter{writer: os.Stdout}, nil
}

// NewTextFileWriter constructor for TextOutputWriter that writes to a file.
func NewTextFileWriter(attributes map[string]any) (OutputWriter, error) {
	path, ok := attributes["path"]
	if !ok {
		return nil, fmt.Errorf("path attribute is required for text_file writer")
	}

	file, err := os.OpenFile(fmt.Sprintf("%s", path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not open log file: %w", err)
	}

	return &TextOutputWriter{writer: file}, nil
}

// MULTI WRITER

// MultiOutputWriter is an OutputWriter that writes logs to multiple writers concurrently.
type MultiOutputWriter struct {
	writers []OutputWriter
}

func (m *MultiOutputWriter) Close() error {
	for _, writer := range m.writers {
		_ = writer.Close()
	}

	return nil
}

// Write sends the log message to all writers concurrently.
func (m *MultiOutputWriter) Write(message LogMessage) {
	var wg sync.WaitGroup

	if len(m.writers) == 0 {
		return
	}

	// If there is only one writer, we can avoid the overhead of goroutines.
	if len(m.writers) == 1 {
		m.writers[0].Write(message)
		return
	}

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
