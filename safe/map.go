package safe

import (
	"encoding/json"
	"sync"
)

// Map ...
type Map[K comparable, V any] struct {
	sync.RWMutex
	cfg *MapConfig
	//
	data map[K]V
	//
	queueForCapacity *Queue[K]
}

// MapConfig ...
type MapConfig struct {
	Capacity int
}

// MapOption ...
type MapOption func(*MapConfig)

// NewMap returns a new safe map
func NewMap[K comparable, V any](opts ...MapOption) *Map[K, V] {
	cfg := &MapConfig{
		Capacity: 0,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	m := &Map[K, V]{
		cfg:  cfg,
		data: make(map[K]V),
	}

	if cfg.Capacity > 0 {
		m.queueForCapacity = NewQueue[K]()
	}

	return m
}

// Set sets a key = value in the map
func (m *Map[K, V]) Set(key K, value V) error {
	m.Lock()
	defer m.Unlock()

	if m.cfg.Capacity > 0 && m.queueForCapacity.Size() >= m.cfg.Capacity {
		keyToRemove := m.queueForCapacity.Dequeue()
		delete(m.data, keyToRemove)
	}

	m.data[key] = value

	if m.cfg.Capacity > 0 {
		m.queueForCapacity.Enqueue(key)
	}

	return nil
}

// Get returns the value for a key
func (m *Map[K, V]) Get(key K) V {
	m.RLock()
	defer m.RUnlock()

	v, ok := m.data[key]
	if !ok {
		return v
	}

	return v
}

// Has returns true if the map contains the key
func (m *Map[K, V]) Has(key K) bool {
	m.RLock()
	defer m.RUnlock()

	if _, ok := m.data[key]; ok {
		return true
	}

	return false
}

// Del deletes a key from the map
func (m *Map[K, V]) Del(key K) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
	}

	return nil
}

// Keys returns all the keys in the map
func (m *Map[K, V]) Keys() []K {
	m.RLock()
	defer m.RUnlock()

	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}

	return keys
}

// Clear removes all elements from the map
func (m *Map[K, V]) Clear() error {
	m.Lock()
	defer m.Unlock()

	m.data = make(map[K]V)

	return nil
}

// ForEach iterates over the map and calls the given function for each key and value
func (m *Map[K, V]) ForEach(f func(K, V) (stop bool)) {
	m.RLock()
	defer m.RUnlock()

	for key, value := range m.data {
		if stop := f(key, value); stop {
			break
		}
	}
}

// Iterator returns a channel that will yield all the keys and values in the map
func (m *Map[K, V]) Iterator() map[K]V {
	return m.ToMap()
}

// ToMap returns the origin map
func (m *Map[K, V]) ToMap() map[K]V {
	m.RLock()
	defer m.RUnlock()

	return m.data
}

// String returns the string representation of the map
func (m *Map[K, V]) String() string {
	bytes, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// MarshalJSON returns the JSON encoding of the map
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	m.RLock()
	defer m.RUnlock()

	return json.Marshal(m.data)
}

// UnmarshalJSON decodes the JSON encoding of the map
func (m *Map[K, V]) UnmarshalJSON(data []byte) error {
	m.Lock()
	defer m.Unlock()

	return json.Unmarshal(data, &m.data)
}
