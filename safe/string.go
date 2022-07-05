package safe

import "sync"

type safeString struct {
	sync.RWMutex
	data string
}

// NewString returns a new safe string
func NewString() *safeString {
	return &safeString{}
}

// Set sets the string to the given value
func (i *safeString) Set(value string) {
	i.Lock()
	defer i.Unlock()

	i.data = value
}

// Get returns the string
func (i *safeString) Get() string {
	i.RLock()
	defer i.RUnlock()

	return i.data
}
