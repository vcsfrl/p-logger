package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

func Build(config Config) (*Logger, error) {
	var outputWriters []OutputWriter
	for _, configWriter := range config.Writers {
		constructorFunction, ok := GetOutputWriterConstructor(configWriter.Name)
		if !ok {
			return nil, fmt.Errorf("BUILD LOGGER: unknown writer: %s", configWriter.Name)
		}

		outputWriter, err := constructorFunction(configWriter.Params)
		if err != nil {
			return nil, fmt.Errorf("BUILD LOGGER: could not create writer: %w", err)
		}

		outputWriters = append(outputWriters, outputWriter)
	}

	// If there is only one writer, we can avoid the overhead of MultiOutputWriter.
	if len(outputWriters) == 1 {
		return NewLogger(outputWriters[0]), nil
	}

	if len(outputWriters) > 1 {
		return NewLogger(&MultiOutputWriter{writers: outputWriters}), nil
	}

	return NewLogger(nil), nil
}

func BuildFromJson(jsonFileName string) (*Logger, error) {
	jsonData, err := os.ReadFile(jsonFileName)
	if err != nil {
		return nil, fmt.Errorf("BUILD LOGGER: could not read json file: %w", err)
	}
	var config Config

	if err = json.Unmarshal(jsonData, &config); err != nil {
		return nil, fmt.Errorf("BUILD LOGGER: could not unmarshal json: %w", err)
	}

	return Build(config)
}
