package io

import (
	gio "io"
)

type writer struct {
	gio.Writer

	fn func(b []byte) (n int, err error)
}

func (w *writer) Write(b []byte) (n int, err error) {
	return w.fn(b)
}

// WriterWrapFunc wraps a function into a Writer.
func WriterWrapFunc(fn func(b []byte) (n int, err error)) gio.Writer {
	return &writer{
		fn: fn,
	}
}

type writeCloser struct {
	gio.Writer
	gio.Closer

	write func(b []byte) (n int, err error)
	close func() error
}

func (w *writeCloser) Write(b []byte) (n int, err error) {
	return w.write(b)
}

func (w *writeCloser) Close() error {
	return w.close()
}

// WriteCloserWrapFunc wraps a function into a WriteCloser.
func WriteCloserWrapFunc(
	write func(b []byte) (n int, err error),
	close func() error,
) gio.WriteCloser {
	return &writeCloser{
		write: write,
		close: close,
	}
}
