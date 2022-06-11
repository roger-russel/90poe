package streamer

import "errors"

var ErrContextCancelationCalled = errors.New("context cancellation got called, streamer will stop right now")
var ErrPathToFileIsDirectory = errors.New("path to file is a directory instead of be a file")
