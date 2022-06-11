package streamer

import "errors"

var ErrContextCancelationCalled = errors.New("context cancellation got called, streamer will stop right now")
