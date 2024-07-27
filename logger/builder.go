package logger

import "fmt"

var OutputWriterSetup map[string]OutputWriterConstructor

func init() {
	OutputWriterSetup = map[string]OutputWriterConstructor{
		"text_stdout": NewTextStdoutWriter,
		"text_file":   NewTextFileWriter,
	}
}

type ConfigWriter struct {
	Name       string
	Attributes map[string]any
}

type Config struct {
	Writers []ConfigWriter
}

func Build(config Config) (*Logger, error) {
	var outputWriters []OutputWriter
	for _, configWriter := range config.Writers {
		constructorFunction, ok := OutputWriterSetup[configWriter.Name]
		if !ok {
			return nil, fmt.Errorf("unknown writer: %s", configWriter.Name)
		}

		outputWriter, err := constructorFunction(configWriter.Attributes)
		if err != nil {
			return nil, fmt.Errorf("could not create writer: %w", err)
		}

		outputWriters = append(outputWriters, outputWriter)
	}

	if len(outputWriters) > 0 {
		return NewLogger(&MultiOutputWriter{writers: outputWriters}), nil
	}

	return NewLogger(nil), nil
}
