package logger

import (
	"bytes"
	"github.com/smarty/assertions/should"
	"testing"

	"github.com/smarty/gunit"
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
	logger.Log("test")

	f.So(output.String(), should.Equal, "test")
}
