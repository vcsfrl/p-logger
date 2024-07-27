package logger

import (
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
)

func TestBuilder(t *testing.T) {
	gunit.Run(new(BuilderFixture), t, gunit.Options.AllSequential())
}

type BuilderFixture struct {
	*gunit.Fixture
}

func (f *BuilderFixture) TestBuildFromConfig() {
	config := Config{}

	logger, err := Build(config)
	//defer logger.Close()

	f.So(logger, should.NotBeNil)
	f.So(err, should.BeNil)
}
