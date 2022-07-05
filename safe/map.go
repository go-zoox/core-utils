package safe

import "sync"

type safeMap struct {
	sync.RWMutex
	data map[string]interface{}
}

// NewMap returns a new safe map
func NewMap() *safeMap {
	return &safeMap{
		data: make(map[string]interface{}),
	}
}

// Set sets a key = value in the map
func (m *safeMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	m.data[key] = value
}

// Get returns the value for a key
func (m *safeMap) Get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.data[key]
}

// Del deletes a key from the map
func (m *safeMap) Del(key string) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
	}
}

// Keys returns all the keys in the map
func (m *safeMap) Keys() []string {
	m.RLock()
	defer m.RUnlock()

	keys := make([]string, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}

	return keys
}

// Clear removes all elements from the map
func (m *safeMap) Clear() {
	m.Lock()
	defer m.Unlock()

	m.data = make(map[string]interface{})
}
