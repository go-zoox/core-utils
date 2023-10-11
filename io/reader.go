package io

import "io"

type reader struct {
	io.Reader
	fn func(p []byte) (n int, err error)
}

// ReaderWrapFunc wraps a function into a Reader.
func ReaderWrapFunc(fn func(p []byte) (n int, err error)) io.Reader {
	return &reader{
		fn: fn,
	}
}

type readerCloser struct {
	io.Reader
	io.Closer

	read  func(p []byte) (n int, err error)
	close func() error
}

func (r *readerCloser) Read(p []byte) (n int, err error) {
	return r.read(p)
}

func (r *readerCloser) Close() error {
	return r.close()
}

// ReaderCloserWrapFunc wraps a function into a ReaderCloser.
func ReaderCloserWrapFunc(
	read func(p []byte) (n int, err error),
	close func() error,
) io.ReadCloser {
	return &readerCloser{
		read:  read,
		close: close,
	}
}
