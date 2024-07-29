package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Builder struct {
}

func (b *Builder) FromConfig(config Config) (*Logger, error) {
	outputWriters, err := b.getOutputWriters(config)
	if err != nil {
		return nil, err
	}

	outputWriter := b.createMultiOutputWriter(outputWriters)

	// Create the logger with the output writer.
	logger := NewLogger(outputWriter)

	// Set the logger's properties based on the configuration.
	if minSeverity, isSet := severityMap[strings.ToUpper(config.MinSeverity)]; isSet {
		logger.MinSeverity = minSeverity
	}
	logger.DefaultTags = config.DefaultTags

	return logger, nil

}

// createMultiOutputWriter creates a single output writer based on the provided list of output writers.
func (b *Builder) createMultiOutputWriter(outputWriters []OutputWriter) OutputWriter {
	var outputWriter OutputWriter

	// If there is only one writer, we can avoid the overhead of MultiOutputWriter.
	if len(outputWriters) == 1 {
		outputWriter = outputWriters[0]
	}

	// If there are multiple writers, we need to use MultiOutputWriter.
	if len(outputWriters) > 1 {
		outputWriter = &MultiOutputWriter{writers: outputWriters}
	}

	return outputWriter
}

func (b *Builder) getOutputWriters(config Config) ([]OutputWriter, error) {
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

	return outputWriters, nil
}

func (b *Builder) FromJson(jsonFileName string) (*Logger, error) {
	jsonData, err := os.ReadFile(jsonFileName)
	if err != nil {
		return nil, fmt.Errorf("BUILD LOGGER: could not read json file: %w", err)
	}
	var config Config

	if err = json.Unmarshal(jsonData, &config); err != nil {
		return nil, fmt.Errorf("BUILD LOGGER: could not unmarshal json: %w", err)
	}

	return b.FromConfig(config)
}

func (b *Builder) LeveledFromJson(jsonFileName string) (*LevelLogger, error) {
	logger, err := b.FromJson(jsonFileName)
	if err != nil {
		return nil, err
	}

	return NewLevelLogger(logger), nil

}

// SOME HELPER FUNCTIONS

// BuildFromConfig creates a new logger based on the provided configuration.
func BuildFromConfig(config Config) (*Logger, error) {
	return new(Builder).FromConfig(config)
}

// BuildFromJson creates a new logger based on the provided JSON configuration file.
func BuildFromJson(jsonFileName string) (*Logger, error) {
	return new(Builder).FromJson(jsonFileName)
}

// BuildLeveledFromJson creates a new leveled logger based on the provided JSON configuration file.
func BuildLeveledFromJson(jsonFileName string) (*LevelLogger, error) {
	return new(Builder).LeveledFromJson(jsonFileName)
}
