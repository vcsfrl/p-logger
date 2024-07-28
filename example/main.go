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

	pLogLevelConfig, _ := logger.BuildLeveledFromJson("./example/config.json")
	defer pLogLevelConfig.Close()
	pLogLevelConfig.Info("Info message - from config")
	pLogLevelConfig.Transaction("11-22-33", map[string]string{"key1": "value1"})
	content, _ := os.ReadFile("./var/log/test.log")
	fmt.Println("Logfile content:")
	fmt.Println(string(content))

	// Output:
	// 2024-07-28T16:33:39+03:00 :: INFO :: Default setup! :: [key1:value1] :: [tag1]
	// 2024-07-28T16:33:39+03:00 :: INFO :: Info message :: [] :: []
	// 2024-07-28T16:33:39+03:00 :: WARN :: Warn message :: [] :: []
	// 2024-07-28T16:33:39+03:00 :: ERROR :: Error message :: [] :: []
	// 2024-07-28T16:33:39+03:00 :: INFO :: Transaction :: [key1:value1 transaction_id:123-ert] :: []
	// 2024-07-28T16:33:39+03:00 :: INFO :: Info message - from config :: [] :: [tag1 tag2 tag3]
	// {"Timestamp":"2024-07-28T16:33:39+03:00","Severity":2,"Content":"Info message - from config","Attributes":null,"Tags":["tag1","tag2","tag3"]}
	// {"Timestamp":"2024-07-28T16:33:39+03:00","Severity":2,"Content":"Transaction","Attributes":{"key1":"value1","transaction_id":"11-22-33"},"Tags":["tag1","tag2","tag3"]}
	// 2024-07-28T16:33:39+03:00 :: INFO :: Transaction :: [key1:value1 transaction_id:11-22-33] :: [tag1 tag2 tag3]
	// Logfile content:
	// 2024-07-28T16:33:39+03:00 :: INFO :: Info message - from config :: [] :: [tag1 tag2 tag3]
	// 2024-07-28T16:33:39+03:00 :: INFO :: Transaction :: [key1:value1 transaction_id:11-22-33] :: [tag1 tag2 tag3]
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
