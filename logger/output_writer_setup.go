package logger

// outputWriterSetup is global variable to the package that holds a map of output writer constructors.
var outputWriterSetup = map[string]OutputWriterConstructor{}

// RegisterOutputWriter registers a new output writer constructor.
// This can be used to add new output writers to the logger from outside this package as well.
func RegisterOutputWriter(name string, constructor OutputWriterConstructor) {
	outputWriterSetup[name] = constructor
}
