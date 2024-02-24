package testutil

import "io"

var _ io.Reader = ReaderEOF{}

// ReaderEOF is an implementation of io.Reader whose Read method always returns
// io.EOF.
type ReaderEOF struct{}

func (ReaderEOF) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}
