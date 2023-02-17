package safe

import "sync"

// String ...
type String struct {
	sync.RWMutex
	data string
}

// NewString returns a new safe string
func NewString() *String {
	return &String{}
}

// Set sets the string to the given value
func (i *String) Set(value string) {
	i.Lock()
	defer i.Unlock()

	i.data = value
}

// Get returns the string
func (i *String) Get() string {
	i.RLock()
	defer i.RUnlock()

	return i.data
}
