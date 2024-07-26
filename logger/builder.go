package logger

type Config struct {
}

func BuildFromConfig(config Config) *Logger {

	return NewLogger(nil)
}
