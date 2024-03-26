package safe

import (
	"encoding/json"
	"sync"

	"github.com/go-zoox/core-utils/cast"
)

// Int64 ...
type Int64 struct {
	sync.RWMutex
	data int64
}

// NewInt64 returns a new safe int
func NewInt64() *Int64 {
	return &Int64{}
}

// Set sets the int to the given value
func (i *Int64) Set(value int64) {
	i.Lock()
	defer i.Unlock()

	i.data = value
}

// Get returns the int
func (i *Int64) Get() int64 {
	i.RLock()
	defer i.RUnlock()

	return i.data
}

// Inc increments the int by step
func (i *Int64) Inc(step int64) {
	i.Lock()
	defer i.Unlock()

	i.data += step
}

// Dec decrements the int by step
func (i *Int64) Dec(step int64) {
	i.Lock()
	defer i.Unlock()

	i.data -= step
}

// String returns the string representation of the int
func (i *Int64) String() string {
	i.RLock()
	defer i.RUnlock()

	return cast.ToString(i.data)
}

// MarshalJSON returns the JSON encoding of the int
func (i *Int64) MarshalJSON() ([]byte, error) {
	i.RLock()
	defer i.RUnlock()

	return json.Marshal(i.data)
}

// UnmarshalJSON decodes the JSON-encoded data and stores the result in the int
func (i *Int64) UnmarshalJSON(data []byte) error {
	i.Lock()
	defer i.Unlock()

	return json.Unmarshal(data, &i.data)
}
