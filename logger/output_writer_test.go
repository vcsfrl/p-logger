package logger

import (
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
)

func TestMultiOutputWriter(t *testing.T) {
	gunit.Run(new(MultiOutputWriterFixture), t)
}

type MultiOutputWriterFixture struct {
	*gunit.Fixture
}

func (f *MultiOutputWriterFixture) TestWriteSingle() {
	buffer1 := new(MemoryWriter)

	writer1 := &TextOutputWriter{writer: buffer1}

	multiWriter := NewMultiOutputWriter(writer1)

	multiWriter.Write(LogMessage{
		Timestamp: mockNow()(),
		Severity:  SeverityDebug,
		Message: Message{
			Content: "example message",
		},
	})

	defer multiWriter.Close()

	f.So(buffer1.String(), should.Equal, "2024-05-01T03:12:03Z :: DEBUG :: example message :: [] :: [] \n")
}

func (f *MultiOutputWriterFixture) TestWriteTwo() {
	buffer1 := new(MemoryWriter)
	buffer2 := new(MemoryWriter)

	writer1 := &TextOutputWriter{writer: buffer1}
	writer2 := &TextOutputWriter{writer: buffer2}

	multiWriter := NewMultiOutputWriter(writer1, writer2)

	multiWriter.Write(LogMessage{
		Timestamp: mockNow()(),
		Severity:  SeverityDebug,
		Message: Message{
			Content: "example message",
		},
	})

	defer multiWriter.Close()

	f.So(buffer1.String(), should.Equal, "2024-05-01T03:12:03Z :: DEBUG :: example message :: [] :: [] \n")
	f.So(buffer2.String(), should.Equal, "2024-05-01T03:12:03Z :: DEBUG :: example message :: [] :: [] \n")
}
