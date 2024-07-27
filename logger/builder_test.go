package logger

import (
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"os"
	"reflect"
	"testing"
)

func TestBuilder(t *testing.T) {
	gunit.Run(new(BuilderFixture), t, gunit.Options.AllSequential())
}

type BuilderFixture struct {
	*gunit.Fixture
	originalStdout *os.File
}

func (f *BuilderFixture) Setup() {
	f.originalStdout = os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
}

func (f *BuilderFixture) Teardown() {
	os.Stdout = f.originalStdout
}

func (f *BuilderFixture) TestBuildEmptyConfig() {
	config := Config{}

	logger, err := Build(config)
	defer logger.Close()

	f.So(logger, should.NotBeNil)
	writer := logger.outputWriter.(*TextOutputWriter)

	f.So(logger, should.NotBeNil)
	f.So(err, should.BeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.TextOutputWriter")
	f.So(writer.writer, should.Equal, os.Stdout)
}
