package streamer

import "errors"

var ErrContextCancelationCalled error = errors.New("context cancellation got called, streamer will stop right now")
