package streaming

import "context"

type Inter interface{}

type Streaming struct{}

func New(ctx context.Context) *Streaming {
	return &Streaming{}
}
