package safe

import (
	"encoding/json"
	"sync"
)

// Stack ...
type Stack[V any] struct {
	sync.RWMutex
	data []V
}

// NewStack returns a new safe stack
func NewStack[V any]() *Stack[V] {
	return &Stack[V]{
		data: make([]V, 0),
	}
}

// Push adds an element to the end of the stack
func (m *Stack[V]) Push(key string, value V) {
	m.Lock()
	defer m.Unlock()

	m.data = append(m.data, value)
}

// Pop removes and returns the last element of the stack
func (m *Stack[V]) Pop(key string) V {
	m.Lock()
	defer m.Unlock()

	if len(m.data) == 0 {
		var v V
		return v
	}

	value := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return value
}

// Size returns the number of elements in the stack
func (m *Stack[V]) Size(key string) int {
	m.RLock()
	defer m.RUnlock()

	return len(m.data)
}

// IsEmpty returns true if the stack is empty
func (m *Stack[V]) IsEmpty(key string) bool {
	return m.Size(key) == 0
}

// Clear removes all elements from the stack
func (m *Stack[V]) Clear(key string) {
	m.Lock()
	defer m.Unlock()

	m.data = make([]V, 0)
}

// Peek returns the last element of the stack
func (m *Stack[V]) Peek(key string) V {
	m.RLock()
	defer m.RUnlock()

	if len(m.data) == 0 {
		var v V
		return v
	}

	return m.data[len(m.data)-1]
}

// ForEach iterates over the stack, calling the given function for each element from top to bottom
func (m *Stack[V]) ForEach(key string, fn func(value V, index int) (stop bool)) {
	m.RLock()
	defer m.RUnlock()

	for i, value := range m.data {
		if stop := fn(value, i); stop {
			break
		}
	}
}

// ForEachReverse iterates over the stack, calling the given function for each element from bottom to top
func (m *Stack[V]) ForEachReverse(key string, fn func(value V, index int) (stop bool)) {
	m.RLock()
	defer m.RUnlock()

	j := 0
	for i := len(m.data) - 1; i >= 0; i-- {
		if stop := fn(m.data[i], j); stop {
			break
		}
		j++
	}
}

// String returns a string representation of the stack
func (m *Stack[V]) String() string {
	bytes, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// MarshalJSON returns the JSON encoding of the stack
func (m *Stack[V]) MarshalJSON() ([]byte, error) {
	m.RLock()
	defer m.RUnlock()

	return json.Marshal(m.data)
}

// UnmarshalJSON decodes the JSON encoding of the stack
func (m *Stack[V]) UnmarshalJSON(data []byte) error {
	var values []V
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	m.data = values

	return nil
}

// ToSlice returns the origin stack
func (m *Stack[V]) ToSlice(key string) []V {
	m.RLock()
	defer m.RUnlock()

	return m.data
}
