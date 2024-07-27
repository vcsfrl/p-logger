package logger

type ConfigWriter struct {
	Name       string
	Attributes map[string]any
}

type Config struct {
	Writers []ConfigWriter
}
