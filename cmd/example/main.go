package main

import (
	"encoding/json"
	"fmt"
	"os"
	"plentysystems-logger/logger"
)

func init() {
	// Register output writer without changes to logger lib.
	logger.RegisterOutputWriter("json_stdout", NewJsonOutputWriter)
}

func main() {
	pLog := logger.NewLogger(nil)
	pLog.Log(
		logger.SeverityInfo,
		logger.Message{
			Content:    "Default setup!",
			Attributes: map[string]string{"key1": "value1"},
			Tags:       []string{"tag1"},
			// Skipped because of min_severity defaults to INFO.
		},
	)

	pLogLevel := logger.NewLevelLogger(pLog)
	defer pLogLevel.Close()
	pLogLevel.Info("Info message")
	pLogLevel.Debug("Debug message")
	pLogLevel.Warn("Warn message")
	pLogLevel.Error("Error message")
	pLogLevel.Transaction("123-ert", map[string]string{"key1": "value1"})

	pLogLevelConfig, err := logger.BuildLeveledFromJson("./cmd/example/config.json")
	fmt.Println(err)
	pLogLevelConfig.Info("Info message")
	pLogLevelConfig.Transaction("11-22-33", map[string]string{"key1": "value1"})
}

type JsonOutputWriter struct {
	encoder *json.Encoder
}

func NewJsonOutputWriter(params map[string]any) (logger.OutputWriter, error) {
	return &JsonOutputWriter{encoder: json.NewEncoder(os.Stdout)}, nil
}

func (j *JsonOutputWriter) Write(message logger.LogMessage) {
	_ = j.encoder.Encode(message)
}

func (j *JsonOutputWriter) Close() error {
	return nil
}
