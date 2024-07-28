package logger

import (
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
)

func TestLevelLogger(t *testing.T) {
	gunit.Run(new(LevelLoggerFixture), t)
}

type LevelLoggerFixture struct {
	*gunit.Fixture
	memoryWriter *MemoryWriter
	logger       *LevelLogger
}

func (f *LevelLoggerFixture) Setup() {
	f.memoryWriter = new(MemoryWriter)
	f.logger = NewLevelLogger(NewLogger(&TextOutputWriter{writer: f.memoryWriter}))
	f.logger.DefaultTags = []string{"tag1", "tag2"}
	f.logger.MinSeverity = SeverityDebug
	// Mock the time.Now function
	f.logger.now = mockNow()
}

func (f *LevelLoggerFixture) Teardown() {
	f.memoryWriter.Reset()
	_ = f.logger.Close()
}

func (f *LevelLoggerFixture) TestLevels() {
	f.logger.Info("test")
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test :: [] :: [tag1 tag2] \n")
}

func (f *LevelLoggerFixture) TestDebug() {
	f.logger.Debug("test")
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: DEBUG :: test :: [] :: [tag1 tag2] \n")
}

func (f *LevelLoggerFixture) TestWarn() {
	f.logger.Warn("test")
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: WARN :: test :: [] :: [tag1 tag2] \n")
}

func (f *LevelLoggerFixture) TestError() {
	f.logger.Error("test")
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: ERROR :: test :: [] :: [tag1 tag2] \n")
}

func (f *LevelLoggerFixture) TestTransaction() {
	f.logger.Transaction("123", map[string]string{"key": "value"})
	f.So(f.memoryWriter.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: Transaction :: [key:value transaction_id:123] :: [tag1 tag2] \n")
}
