package safe

import (
	"sync"
)

// Map ...
type Map struct {
	sync.RWMutex
	cfg *MapConfig
	//
	data map[string]interface{}
	//
	queueForCapacity *Queue
}

// MapConfig ...
type MapConfig struct {
	Capacity int
}

// MapOption ...
type MapOption func(*MapConfig)

// NewMap returns a new safe map
func NewMap(opts ...MapOption) *Map {
	cfg := &MapConfig{
		Capacity: 0,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	m := &Map{
		cfg:  cfg,
		data: make(map[string]interface{}),
	}

	if cfg.Capacity > 0 {
		m.queueForCapacity = NewQueue()
	}

	return m
}

// Set sets a key = value in the map
func (m *Map) Set(key string, value interface{}) error {
	m.Lock()
	defer m.Unlock()

	if m.cfg.Capacity > 0 && m.queueForCapacity.Size() >= m.cfg.Capacity {
		keyToRemove := m.queueForCapacity.Dequeue().(string)
		delete(m.data, keyToRemove)
	}

	m.data[key] = value

	if m.cfg.Capacity > 0 {
		m.queueForCapacity.Enqueue(key)
	}

	return nil
}

// Get returns the value for a key
func (m *Map) Get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()

	v, ok := m.data[key]
	if !ok {
		return nil
	}

	return v
}

// Has returns true if the map contains the key
func (m *Map) Has(key string) bool {
	m.RLock()
	defer m.RUnlock()

	if _, ok := m.data[key]; ok {
		return true
	}

	return false
}

// Del deletes a key from the map
func (m *Map) Del(key string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
	}

	return nil
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
func (m *Map) Clear() error {
	m.Lock()
	defer m.Unlock()

	m.data = make(map[string]interface{})

	return nil
}

// Iterator returns a channel that will yield all the keys and values in the map
func (m *Map) Iterator() map[string]interface{} {
	return m.ToMap()
}

// ToMap returns the origin map
func (m *Map) ToMap() map[string]interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.data
}
