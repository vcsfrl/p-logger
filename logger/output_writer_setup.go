package logger

import "sync"

// outputWriterSetup is global variable to the package that holds a map of output writer constructors.
var outputWriterSetup = map[string]OutputWriterConstructor{}
var outputWriterSetupLock = &sync.RWMutex{}

func init() {
	RegisterOutputWriter("text_stdout", NewTextStdoutWriter)
	RegisterOutputWriter("text_file", NewTextFileWriter)
}

// RegisterOutputWriter registers a new output writer constructor.
// This can be used to add new output writers to the logger from outside this package as well.
func RegisterOutputWriter(name string, constructor OutputWriterConstructor) {
	outputWriterSetupLock.Lock()
	defer outputWriterSetupLock.Unlock()
	outputWriterSetup[name] = constructor
}

func GetOutputWriterConstructor(name string) (OutputWriterConstructor, bool) {
	outputWriterSetupLock.RLock()
	defer outputWriterSetupLock.RUnlock()
	value, exists := outputWriterSetup[name]

	return value, exists

}
