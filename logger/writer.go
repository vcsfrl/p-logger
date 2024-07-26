package logger

type OutputWriter interface {
	Write(message LogMessage)
}
