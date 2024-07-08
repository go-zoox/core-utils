package strings

import (
	gs "strings"
	"sync"
)

type Builder interface {
	String() string
	Len() int
	Write(b []byte) (int, error)
	Reset()
}

type builder struct {
	core *gs.Builder
	sync.RWMutex
}

func NewBuilder() Builder {
	return &builder{
		core: &gs.Builder{},
	}
}

func (b *builder) String() string {
	b.RLock()
	defer b.RUnlock()

	return b.core.String()
}

func (b *builder) Len() int {
	b.RLock()
	defer b.RUnlock()

	return b.core.Len()
}

func (b *builder) Write(p []byte) (int, error) {
	b.Lock()
	defer b.Unlock()

	return b.core.Write(p)
}

func (b *builder) Reset() {
	b.Lock()
	defer b.Unlock()

	b.core.Reset()
}
