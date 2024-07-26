package logger

import (
	"bytes"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
)

func TestMultiOutputWriter(t *testing.T) {
	gunit.Run(new(MultiOutputWriterFixture), t, gunit.Options.AllSequential())
}

type MultiOutputWriterFixture struct {
	*gunit.Fixture
}

func (f *MultiOutputWriterFixture) TestWrite() {
	buffer1 := new(bytes.Buffer)
	buffer2 := new(bytes.Buffer)

	writer1 := &TextOutputWriter{writer: buffer1}
	writer2 := &TextOutputWriter{writer: buffer2}

	multiWriter := NewMultiOutputWriter(writer1, writer2)

	multiWriter.Write(LogMessage{
		Timestamp: mockNow()(),
		Severity:  SeverityInfo,
		Message: Message{
			Content: "test message",
		},
	})

	f.So(buffer1.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test message :: [] :: [] \n")
	f.So(buffer2.String(), should.Equal, "2024-05-01T03:12:03Z :: INFO :: test message :: [] :: [] \n")
}
