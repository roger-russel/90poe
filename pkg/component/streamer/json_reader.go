package streamer

type readerStage uint8

const (
	rdrStageFindStart           readerStage = 0
	rdrStageFindKeyNameStart    readerStage = 1
	rdrStageReadingKeyName      readerStage = 2
	rdrStageFindKeyContentStart readerStage = 3
	rdrStageReadingKeyContent   readerStage = 4
)

const (
	tokenBackslash         byte = 92  // \
	tokenLeftCurlyBracket  byte = 123 // {
	tokenRightCurlyBracket byte = 125 // }
	tokenDoubleQuote       byte = 34  // "
)

func (j *JSON) handleBackslash(c byte) bool {
	switch {
	case c == tokenBackslash && !j.controlEscapedChar:
		j.controlEscapedChar = true
		j.escapedChar = false

		return true
	case j.controlEscapedChar:
		j.escapedChar = true
		j.controlEscapedChar = false

		return false
	default:
		j.controlEscapedChar = false
		j.escapedChar = false

		return false
	}
}

func (j *JSON) reader(buf []byte) {
	switch j.readerStage {
	case rdrStageFindStart:
		j.readerStart(buf)
	case rdrStageFindKeyNameStart:
		j.readerFindKeyNameStart(buf)
	case rdrStageReadingKeyName:
		j.readerReadingKeyName(buf)
	case rdrStageFindKeyContentStart:
		j.readerFindKeyContentStart(buf)
	case rdrStageReadingKeyContent:
		j.readerKeyContent(buf)
	}
}

func (j *JSON) readerStart(buf []byte) {
	for i, b := range buf {
		if b == tokenLeftCurlyBracket {
			j.readerStage = rdrStageFindKeyNameStart
			j.readerFindKeyNameStart(buf[i:])

			break
		}
	}
}

func (j *JSON) readerFindKeyNameStart(buf []byte) {
	for i, b := range buf {
		if j.readerStage == rdrStageFindKeyNameStart && b == tokenDoubleQuote {
			j.readerStage = rdrStageReadingKeyName
			if len(buf) > i+1 {
				j.readerReadingKeyName(buf[i+1:])
			}

			break
		}
	}
}

func (j *JSON) readerReadingKeyName(buf []byte) {
	for i, b := range buf {
		if j.readerStage == rdrStageReadingKeyName {
			_ = j.handleBackslash(b)

			if b == tokenDoubleQuote && !j.escapedChar {
				j.readerStage = rdrStageFindKeyContentStart
				j.readerFindKeyContentStart(buf[i:])

				break
			}

			j.bufferKeyName = append(j.bufferKeyName, b)
		}
	}
}

func (j *JSON) readerFindKeyContentStart(buf []byte) {
	for i, b := range buf {
		if b == tokenLeftCurlyBracket {
			j.readerStage = rdrStageReadingKeyContent
			j.lookingToken = tokenRightCurlyBracket
			j.lifoToken = append(j.lifoToken, tokenRightCurlyBracket)
			j.readerKeyContent(buf[i:])

			break
		}
	}
}

func (j *JSON) readerKeyContent(buf []byte) {
	for i, b := range buf {
		j.bufferKeyContent = append(j.bufferKeyContent, b)
		_ = j.handleBackslash(b)
		if b == j.lookingToken && !j.escapedChar {
			if len(j.lifoToken) > 1 {
				j.lifoToken = j.lifoToken[:len(j.lifoToken)-1]
				j.lookingToken = j.lifoToken[len(j.lifoToken)-1]
			} else {
				j.sendJSON()
				j.readerStage = rdrStageFindKeyNameStart
				j.readerFindKeyNameStart(buf[i:])

				break
			}
		}
	}
}
