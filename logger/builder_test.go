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
	builder        *Builder
}

func (f *BuilderFixture) Setup() {
	// Mock stdout.
	f.originalStdout = os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
	f.builder = new(Builder)
}

func (f *BuilderFixture) Teardown() {
	os.Stdout = f.originalStdout
}

func (f *BuilderFixture) TestBuildEmptyConfig() {
	config := Config{}

	logger, err := f.builder.FromConfig(config)
	defer func() {
		err := logger.Close()
		f.So(err, should.BeNil)
	}()

	f.So(logger, should.NotBeNil)
	if logger == nil {
		return
	}
	writer := logger.outputWriter.(*TextOutputWriter)

	f.So(logger, should.NotBeNil)
	f.So(err, should.BeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.TextOutputWriter")
	f.So(writer.writer, should.Equal, os.Stdout)
}

func (f *BuilderFixture) TestBuildOneOutputWriter() {
	config := Config{
		Writers: []ConfigWriter{
			{
				Name: "text_file",
				Params: map[string]interface{}{
					"path": "./../var/log/example.log",
				},
			},
		},
	}

	logger, err := f.builder.FromConfig(config)
	defer func() {
		err := logger.Close()
		f.So(err, should.BeNil)
	}()

	f.So(logger, should.NotBeNil)

	f.So(logger, should.NotBeNil)
	f.So(err, should.BeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.TextOutputWriter")
	writer := logger.outputWriter.(*TextOutputWriter)
	f.So(reflect.TypeOf(writer.writer).String(), should.Equal, "*os.File")
}

func (f *BuilderFixture) TestBuildTwoOutputWriters() {
	config := Config{
		Writers: []ConfigWriter{
			{
				Name: "text_file",
				Params: map[string]interface{}{
					"path": "./../var/log/example.log",
				},
			}, {
				Name: "text_stdout",
			},
		},
	}

	logger, err := f.builder.FromConfig(config)
	defer func() {
		err := logger.Close()
		f.So(err, should.BeNil)
	}()

	f.So(logger, should.NotBeNil)
	f.So(err, should.BeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.MultiOutputWriter")
	multiOutputWriter := logger.outputWriter.(*MultiOutputWriter)

	fileWriter := multiOutputWriter.writers[0].(*TextOutputWriter)
	f.So(reflect.TypeOf(fileWriter.writer).String(), should.Equal, "*os.File")
	f.So(fileWriter.writer, should.NotEqual, os.Stdout)

	stdoutWriter := multiOutputWriter.writers[1].(*TextOutputWriter)
	f.So(stdoutWriter.writer, should.Equal, os.Stdout)
}

func (f *BuilderFixture) TestBuildMinSeverityAndDefaultTags() {
	config := Config{
		Writers: []ConfigWriter{
			{
				Name: "text_file",
				Params: map[string]interface{}{
					"path": "./../var/log/example.log",
				},
			}, {
				Name: "text_stdout",
			},
		},
		MinSeverity: "WARN",
		DefaultTags: []string{"tag1", "tag2"},
	}

	logger, err := f.builder.FromConfig(config)
	defer func() {
		err := logger.Close()
		f.So(err, should.BeNil)
	}()

	f.So(logger, should.NotBeNil)
	f.So(err, should.BeNil)
	f.So(logger.MinSeverity, should.Equal, SeverityWarning)
	f.So(logger.DefaultTags, should.Resemble, []string{"tag1", "tag2"})
}

func (f *BuilderFixture) TestBuildFromJson() {
	logger, err := f.builder.FromJson("testdata/example_config_valid.json")
	defer func() {
		err := logger.Close()
		f.So(err, should.BeNil)
	}()

	f.So(err, should.BeNil)
	f.So(logger, should.NotBeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.MultiOutputWriter")
	multiOutputWriter := logger.outputWriter.(*MultiOutputWriter)

	fileWriter := multiOutputWriter.writers[0].(*TextOutputWriter)
	f.So(reflect.TypeOf(fileWriter.writer).String(), should.Equal, "*os.File")
	f.So(fileWriter.writer, should.NotEqual, os.Stdout)

	stdoutWriter := multiOutputWriter.writers[1].(*TextOutputWriter)
	f.So(stdoutWriter.writer, should.Equal, os.Stdout)

	f.So(logger.MinSeverity, should.Equal, SeverityWarning)
	f.So(logger.DefaultTags, should.Resemble, []string{"tag1", "tag2", "tag3"})
}

func (f *BuilderFixture) TestBuildLeveledFromJson() {
	logger, err := f.builder.LeveledFromJson("testdata/example_config_valid.json")
	defer func() {
		err := logger.Close()
		f.So(err, should.BeNil)
	}()

	f.So(err, should.BeNil)
	f.So(logger, should.NotBeNil)
	f.So(reflect.TypeOf(logger.outputWriter).String(), should.Equal, "*logger.MultiOutputWriter")
	multiOutputWriter := logger.outputWriter.(*MultiOutputWriter)

	fileWriter := multiOutputWriter.writers[0].(*TextOutputWriter)
	f.So(reflect.TypeOf(fileWriter.writer).String(), should.Equal, "*os.File")
	f.So(fileWriter.writer, should.NotEqual, os.Stdout)

	stdoutWriter := multiOutputWriter.writers[1].(*TextOutputWriter)
	f.So(stdoutWriter.writer, should.Equal, os.Stdout)

	f.So(logger.MinSeverity, should.Equal, SeverityWarning)
	f.So(logger.DefaultTags, should.Resemble, []string{"tag1", "tag2", "tag3"})
}

func (f *BuilderFixture) TestBuildFromJsonInvalidData() {
	logger, err := f.builder.FromJson("testdata/example_config_invalid.json")

	f.So(err, should.NotBeNil)
	f.So(logger, should.BeNil)

	logger, err = f.builder.FromJson("testdata/example_config_not_exists.json")

	f.So(err, should.NotBeNil)
	f.So(logger, should.BeNil)
}
