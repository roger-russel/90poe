package streamer

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

var _ Inter = (*JSON)(nil)

const DefaultBufferParserSize int = 4096 // 4kb

type JSONConfig struct {
	BufferParserSize int
}

type JSON struct {
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

func NewJSON(conf JSONConfig) *JSON {
	bufferParserSize := conf.BufferParserSize

	if bufferParserSize < 1 {
		bufferParserSize = DefaultBufferParserSize
	}

	return &JSON{
		bufferParserSize: bufferParserSize,
		readerStage:      rdrStageFindStart,
	}
}

func (j *JSON) Stream(ctx context.Context, reader io.Reader, chDataOutput chan Data) error {
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

func (j *JSON) StreamFile(ctx context.Context, filePath string, chDataOutput chan Data) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", filePath, err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting stat of file %s: %w", filePath, err)
	}

	if fileInfo.IsDir() {
		return ErrPathToFileIsDirectory
	}

	log.Printf("open file %v to streaming", filePath)

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing file", filePath, err)
		}
	}()

	return j.Stream(ctx, file, chDataOutput)
}

func (j *JSON) sendJSON() {
	j.chDataOutput <- Data{
		KeyName:    j.bufferKeyName,
		KeyContent: j.bufferKeyContent,
	}

	j.bufferKeyContent = []byte{}
	j.bufferKeyName = []byte{}
	j.lifoToken = []byte{}
}
