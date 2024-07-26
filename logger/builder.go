package logger

type ConfigWriter struct {
	Name       string
	Attributes map[string]string
}

type Config struct {
	Writers []ConfigWriter
}

func BuildFromConfig(config Config) (*Logger, error) {

	for _, configWriter := range config.Writers {
		if configWriter.Name == "text" {

		}
	}

	return NewLogger(nil), nil
}
