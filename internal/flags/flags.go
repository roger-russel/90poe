package flags

import (
	"flag"
)

type Flags struct {
	ParserBufferSize int
	File             string
}

func Load() Flags {
	f := flag.String("file", "", "read file on path")
	s := flag.Int("buffer-size", 0, "parser buffer size")
	flag.Parse()

	fl := Flags{
		File:             *f,
		ParserBufferSize: *s,
	}

	return fl
}
