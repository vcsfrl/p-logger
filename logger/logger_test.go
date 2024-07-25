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

	// Mock the logger
	logger.now = mockNow()
	logger.Log("test")

	f.So(output.String(), should.Equal, "2024-05-01T03:12:03Z : test")
}

// Mock the time.Now function
func mockNow() func() time.Time {
	return func() time.Time {
		return time.Date(2024, 5, 1, 3, 12, 3, 0, time.UTC)
	}
}
