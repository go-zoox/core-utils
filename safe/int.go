package safe

import "sync"

// Int ...
type Int struct {
	sync.RWMutex
	data int
}

// NewInt returns a new safe int
func NewInt() *Int {
	return &Int{}
}

// Set sets the int to the given value
func (i *Int) Set(value int) {
	i.Lock()
	defer i.Unlock()

	i.data = value
}

// Get returns the int
func (i *Int) Get() int {
	i.RLock()
	defer i.RUnlock()

	return i.data
}

// Inc increments the int by step
func (i *Int) Inc(step int) {
	i.Lock()
	defer i.Unlock()

	i.data += step
}

// Dec decrements the int by step
func (i *Int) Dec(step int) {
	i.Lock()
	defer i.Unlock()

	i.data -= step
}
