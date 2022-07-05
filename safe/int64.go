package safe

import "sync"

type safeInt64 struct {
	sync.RWMutex
	data int64
}

// NewInt returns a new safe int
func NewInt64() *safeInt64 {
	return &safeInt64{}
}

// Set sets the int to the given value
func (i *safeInt64) Set(value int64) {
	i.Lock()
	defer i.Unlock()

	i.data = value
}

// Get returns the int
func (i *safeInt64) Get() int64 {
	i.RLock()
	defer i.RUnlock()

	return i.data
}
