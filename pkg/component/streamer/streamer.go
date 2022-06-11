package streamer

import (
	"context"
	"io"
)

type Inter interface {
	Stream(ctx context.Context, reader io.Reader, chDataOutput chan Data) error
	StreamFile(ctx context.Context, filePath string, chDataOutput chan Data) error
}

type Data struct {
	KeyName    []byte
	KeyContent []byte
}

type Message struct {
	Key     []byte
	Content []byte
}
