package streamer

import (
	"context"
	"io"
	"os"
)

const DefaultBufferParserSize int = 4096 // 4kb

type JsonConfig struct {
	BufferParserSize int
}

type Json struct {
	bufferParserSize int
	chDataOutput     chan Data

	lookingToken       byte
	escapedChar        bool
	controlEscapedChar bool
	readerStage        readerStage
	bufferKeyContent   []byte
	bufferKeyName      []byte
	lifoToken          []byte
}

func NewJson(conf JsonConfig) Inter {
	bufferParserSize := conf.BufferParserSize

	if bufferParserSize < 1 {
		bufferParserSize = DefaultBufferParserSize
	}

	return &Json{
		bufferParserSize: bufferParserSize,
		readerStage:      rdrStageFindStart,
	}
}

func (j *Json) Stream(ctx context.Context, reader io.Reader, chDataOutput chan Data) error {
	j.chDataOutput = chDataOutput
	for {
		select {
		case <-ctx.Done():
			return ErrContextCancelationCalled
		default:
			p := make([]byte, j.bufferParserSize)
			_, err := reader.Read(p)
			if err == io.EOF {
				return nil
			}
			j.reader(p)
		}
	}
}

func (j *Json) StreamFile(ctx context.Context, filePath string, chDataOutput chan Data) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	return j.Stream(ctx, file, chDataOutput)
}

func (j *Json) sendJson() {
	j.chDataOutput <- Data{
		KeyName:    j.bufferKeyName,
		KeyContent: j.bufferKeyContent,
	}

	j.bufferKeyContent = []byte{}
	j.bufferKeyName = []byte{}
	j.lifoToken = []byte{}
}
