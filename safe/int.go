package safe

import (
	"encoding/json"
	"sync"

	"github.com/go-zoox/core-utils/cast"
)

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

// String returns the string representation of the int
func (i *Int) String() string {
	i.RLock()
	defer i.RUnlock()

	return cast.ToString(i.data)
}

// MarshalJSON returns the JSON encoding of the int
func (i *Int) MarshalJSON() ([]byte, error) {
	i.RLock()
	defer i.RUnlock()

	return json.Marshal(i.data)
}

// UnmarshalJSON decodes the JSON-encoded data and stores the result in the int
func (i *Int) UnmarshalJSON(data []byte) error {
	i.Lock()
	defer i.Unlock()

	return json.Unmarshal(data, &i.data)
}
