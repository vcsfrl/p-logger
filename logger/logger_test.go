package logger

import (
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"os"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	gunit.Run(new(LoggerFixture), t, gunit.Options.AllSequential())
}

type LoggerFixture struct {
	*gunit.Fixture
	memoryWriter *MemoryWriter
	logger       *Logger
}

func (f *LoggerFixture) Setup() {
	f.memoryWriter = new(MemoryWriter)
	f.logger = NewLogger(&TextOutputWriter{writer: f.memoryWriter})
	// Mock the time.Now function
	f.logger.now = mockNow()
}

func (f *LoggerFixture) Teardown() {
	f.memoryWriter.Reset()
	_ = f.logger.Close()
}

func (f *LoggerFixture) TestConstructorDefaults() {
	logger := NewLogger(nil)
	writer := logger.outputWriter.(*TextOutputWriter)

	// Check if the default writer is os.Stdout
	f.So(logger, should.NotBeNil)
	f.So(logger.outputWriter, should.NotBeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.TextOutputWriter")
	f.So(writer.writer, should.Equal, os.Stdout)
}

func (f *LoggerFixture) TestLog() {
	f.logger.Log(SeverityInfo, Message{
		Content: "example",
		Attributes: map[string]string{
			"attr-key-1": "attr-value1",
			"attr-key-2": "attr-value2",
		},
		Tags: []string{"tag1", "tag2"},
	})

	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: example :: [attr-key-1:attr-value1 attr-key-2:attr-value2] :: [tag1 tag2] \n")
}

func (f *LoggerFixture) TestLogDefaultSeverity() {
	// Test severity default on invalid value.
	f.logger.Log(12, Message{
		Content: "test1",
	})
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test1 :: [] :: [] \n")
}

func (f *LoggerFixture) TestLogMinSeverity() {
	f.logger.MinSeverity = SeverityWarning
	// Severity info should be skipped.
	f.logger.Log(SeverityInfo, Message{
		Content: "test1",
	})
	f.So(f.memoryWriter.String(), should.Equal, "")

	f.logger.Log(SeverityWarning, Message{
		Content: "test2",
	})
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: WARN :: test2 :: [] :: [] \n")
}

func (f *LoggerFixture) TestLogDefaultTags() {
	f.logger.DefaultTags = []string{"tag1", "tag2"}
	// Test severity default on invalid value.
	f.logger.Log(SeverityInfo, Message{
		Content: "test1",
	})
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test1 :: [] :: [tag1 tag2] \n")
}
