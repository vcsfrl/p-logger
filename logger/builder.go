package logger

import "fmt"

func Build(config Config) (*Logger, error) {
	var outputWriters []OutputWriter
	for _, configWriter := range config.Writers {
		constructorFunction, ok := outputWriterSetup[configWriter.Name]
		if !ok {
			return nil, fmt.Errorf("BUILD LOGGER: unknown writer: %s", configWriter.Name)
		}

		outputWriter, err := constructorFunction(configWriter.Attributes)
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
