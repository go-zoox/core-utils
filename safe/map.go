package safe

import "sync"

type Map struct {
	sync.RWMutex
	data map[string]interface{}
}

// NewMap returns a new safe map
func NewMap() *Map {
	return &Map{
		data: make(map[string]interface{}),
	}
}

// Set sets a key = value in the map
func (m *Map) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	m.data[key] = value
}

// Get returns the value for a key
func (m *Map) Get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.data[key]
}

// Has returns true if the map contains the key
func (m *Map) Has(key string) bool {
	m.RLock()
	defer m.RUnlock()

	if _, ok := m.data[key]; ok {
		return true
	} else {
		return false
	}
}

// Del deletes a key from the map
func (m *Map) Del(key string) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
	}
}

// Keys returns all the keys in the map
func (m *Map) Keys() []string {
	m.RLock()
	defer m.RUnlock()

	keys := make([]string, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}

	return keys
}

// Clear removes all elements from the map
func (m *Map) Clear() {
	m.Lock()
	defer m.Unlock()

	m.data = make(map[string]interface{})
}

// Iterator returns a channel that will yield all the keys and values in the map
func (m *Map) Iterator() map[string]interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.data
}
