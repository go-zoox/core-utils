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

type writerCloser struct {
	gio.Writer
	gio.Closer

	write func(b []byte) (n int, err error)
	close func() error
}

func (w *writerCloser) Write(b []byte) (n int, err error) {
	return w.write(b)
}

func (w *writerCloser) Close() error {
	return w.close()
}

// WriterCloserWrapFunc wraps a function into a WriterCloser.
func WriterCloserWrapFunc(
	write func(b []byte) (n int, err error),
	close func() error,
) gio.WriteCloser {
	return &writerCloser{
		write: write,
		close: close,
	}
}
