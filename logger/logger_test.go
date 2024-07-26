package logger

import (
	"bytes"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	gunit.Run(new(LoggerFixture), t, gunit.Options.AllSequential())
}

type LoggerFixture struct {
	*gunit.Fixture
}

func (f *LoggerFixture) Setup() {

}

func (f *LoggerFixture) Teardown() {

}

func (f *LoggerFixture) TestConstructor() {
	logger := NewLogger(nil)

	f.So(logger, should.NotBeNil)
}

func (f *LoggerFixture) TestLog() {
	output := new(bytes.Buffer)
	logger := NewLogger(output)

	// Mock the time.Now function
	logger.now = mockNow()
	logger.Log(SeverityInfo, Message{
		Content: "test",
		Attributes: map[string]string{
			"attr-key-1": "attr-value1",
			"attr-key-2": "attr-value2",
		},
		Tags: []string{"tag1", "tag2"},
	})

	f.So(output.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test :: [attr-key-1:attr-value1 attr-key-2:attr-value2] :: [tag1 tag2] \n")
	output.Reset()

	// Test with default severity used on invalid param
	logger.Log("", Message{
		Content: "test1",
	})
	f.So(output.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test1 :: [] :: [] \n")
}

func (f *LoggerFixture) TestLogDefaultSeverity() {
	output := new(bytes.Buffer)
	logger := NewLogger(output)

	// Mock the time.Now function
	logger.now = mockNow()

	// Test severity default on invalid value.
	logger.Log("", Message{
		Content: "test1",
	})
	f.So(output.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test1 :: [] :: [] \n")
}

// Mocks

func mockNow() func() time.Time {
	return func() time.Time {
		return time.Date(2024, 5, 1, 3, 12, 3, 0, time.UTC)
	}
}
