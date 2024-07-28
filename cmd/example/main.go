package main

import (
	"fmt"
	"io"
	"os"
	"plentysystems-logger/logger"
)

func init() {
	// Register output writer without changes to logger lib.
	logger.RegisterOutputWriter("json_file", NewJsonOutputWriter)
}

func main() {
	pLog := logger.NewLogger(nil)
	pLog.Log(
		logger.SeverityInfo,
		logger.Message{
			Content:    "Default setup!",
			Attributes: map[string]string{"key1": "value1"},
			Tags:       []string{"tag1"},
		},
	)

	pLogLevel := logger.NewLevelLogger(pLog)
	pLogLevel.Info("Info message")
	// Skipped because of min_severity defaults to INFO.
	pLogLevel.Debug("Debug message")
	pLogLevel.Warn("Warn message")
	pLogLevel.Error("Error message")
	pLogLevel.Transaction("123-ert", map[string]string{"key1": "value1"})
	_ = pLogLevel.Close()
}

type JsonOutputWriter struct {
	writer io.WriteCloser
}

func NewJsonOutputWriter(params map[string]any) (logger.OutputWriter, error) {
	file, err := os.OpenFile(fmt.Sprintf("%s", "./../../var/log/example.sjon"), os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("NEW JSON WRITER: could not open log file: %w", err)
	}
	return &JsonOutputWriter{file}, nil
}

func (j *JsonOutputWriter) Write(message logger.LogMessage) {

}

func (j *JsonOutputWriter) Close() error {
	return j.writer.Close()
}
