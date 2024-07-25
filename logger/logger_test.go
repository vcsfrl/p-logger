package logger

import (
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

func (f *LoggerFixture) TestCreateInstance() {
	NewLogger()
}
