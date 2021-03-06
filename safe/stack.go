package safe

import "sync"

type Stack struct {
	sync.RWMutex
	data []interface{}
}

// NewStack returns a new safe stack
func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
	}
}

// Push adds an element to the end of the stack
func (m *Stack) Push(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	m.data = append(m.data, value)
}

// Pop removes and returns the last element of the stack
func (m *Stack) Pop(key string) interface{} {
	m.Lock()
	defer m.Unlock()

	if len(m.data) == 0 {
		return nil
	}

	value := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return value
}

// Size returns the number of elements in the stack
func (m *Stack) Size(key string) int {
	m.RLock()
	defer m.RUnlock()

	return len(m.data)
}

// IsEmpty returns true if the stack is empty
func (m *Stack) IsEmpty(key string) bool {
	return m.Size(key) == 0
}

// Clear removes all elements from the stack
func (m *Stack) Clear(key string) {
	m.Lock()
	defer m.Unlock()

	m.data = make([]interface{}, 0)
}

// Peek returns the last element of the stack
func (m *Stack) Peek(key string) interface{} {
	m.RLock()
	defer m.RUnlock()

	if len(m.data) == 0 {
		return nil
	}

	return m.data[len(m.data)-1]
}
