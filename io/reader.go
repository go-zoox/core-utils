package io

import "io"

type reader struct {
	io.Reader
	read func(p []byte) (n int, err error)
}

func (r *reader) Read(p []byte) (n int, err error) {
	return r.read(p)
}

// ReadWrapFunc wraps a function into a Reader.
func ReadWrapFunc(read func(p []byte) (n int, err error)) io.Reader {
	return &reader{
		read: read,
	}
}

type readCloser struct {
	io.Reader
	io.Closer

	read  func(p []byte) (n int, err error)
	close func() error
}

func (r *readCloser) Read(p []byte) (n int, err error) {
	return r.read(p)
}

func (r *readCloser) Close() error {
	return r.close()
}

// ReadCloserWrapFunc wraps a function into a ReadCloser.
func ReadCloserWrapFunc(
	read func(p []byte) (n int, err error),
	close func() error,
) io.ReadCloser {
	return &readCloser{
		read:  read,
		close: close,
	}
}
