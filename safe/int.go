package safe

import "sync"

type safeInt struct {
	sync.RWMutex
	data int
}

// NewInt returns a new safe int
func NewInt() *safeInt {
	return &safeInt{}
}

// Set sets the int to the given value
func (i *safeInt) Set(value int) {
	i.Lock()
	defer i.Unlock()

	i.data = value
}

// Get returns the int
func (i *safeInt) Get() int {
	i.RLock()
	defer i.RUnlock()

	return i.data
}
