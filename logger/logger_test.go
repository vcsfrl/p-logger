package logger

import (
	"bytes"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
)

func TestLogger(t *testing.T) {
	gunit.Run(new(LoggerFixture), t, gunit.Options.AllSequential())
}

type LoggerFixture struct {
	*gunit.Fixture
	writer *bytes.Buffer
	logger *Logger
}

func (f *LoggerFixture) Setup() {
	f.writer = new(bytes.Buffer)
	f.logger = NewLogger(&TextOutputWriter{writer: f.writer})
	// Mock the time.Now function
	f.logger.now = mockNow()
}

func (f *LoggerFixture) Teardown() {
	f.writer.Reset()
}

func (f *LoggerFixture) TestConstructorDefaults() {
	logger := NewLogger(nil)

	f.So(logger, should.NotBeNil)
	f.So(logger.outputWriter, should.NotBeNil)
}

func (f *LoggerFixture) TestLog() {
	f.logger.Log(SeverityInfo, Message{
		Content: "test",
		Attributes: map[string]string{
			"attr-key-1": "attr-value1",
			"attr-key-2": "attr-value2",
		},
		Tags: []string{"tag1", "tag2"},
	})

	f.So(f.writer.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test :: [attr-key-1:attr-value1 attr-key-2:attr-value2] :: [tag1 tag2] \n")
}

func (f *LoggerFixture) TestLogDefaultSeverity() {
	// Test severity default on invalid value.
	f.logger.Log("", Message{
		Content: "test1",
	})
	f.So(f.writer.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test1 :: [] :: [] \n")
}
