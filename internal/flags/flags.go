package flags

import (
	"flag"
)

type Flags struct {
	ParserBufferSize int
	File             string
}

func Load() Flags {
	return Flags{
		File:             *flag.String("file", "", "file to be read"),
		ParserBufferSize: *flag.Int("parserBufferSize", 0, "parser buffer size"),
	}
}
