package logger

import (
	"bytes"
	"github.com/smarty/gunit"
	"testing"
)

func TestBuilder(t *testing.T) {
	gunit.Run(new(BuilderFixture), t, gunit.Options.AllSequential())
}

type BuilderFixture struct {
	*gunit.Fixture
	writer *bytes.Buffer
}

func (f *BuilderFixture) TestBuildFromConfig() {
	//config := Config{}
	//
	//logger, err := BuildFromConfig(config)
	////defer logger.Close()
	//
	//f.So(logger, should.NotBeNil)
	//f.So(err, should.BeNil)

}
