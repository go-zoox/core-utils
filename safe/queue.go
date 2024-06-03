package safe

import (
	"encoding/json"
	"sync"
)

// Queue ...
type Queue[V any] struct {
	sync.RWMutex
	data []V
	//
	cfg *QueueConfig
}

// QueueConfig ...
type QueueConfig struct {
	Capacity int
}

// QueueOption ...
type QueueOption func(*QueueConfig)

// NewQueue returns a new safe queue
func NewQueue[V any](opts ...QueueOption) *Queue[V] {
	cfg := &QueueConfig{
		Capacity: 0,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return &Queue[V]{
		cfg:  cfg,
		data: make([]V, 0),
	}
}

// Enqueue adds an element to the end of the queue
func (q *Queue[V]) Enqueue(value V) {
	q.Lock()
	defer q.Unlock()

	q.data = append(q.data, value)

	// check capacity when push
	if q.cfg.Capacity > 0 && len(q.data) > q.cfg.Capacity {
		q.data = q.data[len(q.data)-q.cfg.Capacity:]
	}
}

// Dequeue removes and returns the first element of the queue
func (q *Queue[V]) Dequeue() V {
	q.Lock()
	defer q.Unlock()

	if len(q.data) == 0 {
		var v V
		return v
	}

	value := q.data[0]
	q.data = q.data[1:]
	return value
}

// Size returns the number of elements in the queue
func (q *Queue[V]) Size() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.data)
}

// Front returns the first element of the queue
func (q *Queue[V]) Front() V {
	q.RLock()
	defer q.RUnlock()

	if len(q.data) == 0 {
		var v V
		return v
	}

	return q.data[0]
}

// Back returns the last element of the queue
func (q *Queue[V]) Back() V {
	q.RLock()
	defer q.RUnlock()

	if len(q.data) == 0 {
		var v V
		return v
	}

	return q.data[len(q.data)-1]
}

// IsEmpty returns true if the queue is empty
func (q *Queue[V]) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.data) == 0
}

// Clear removes all elements from the queue
func (q *Queue[V]) Clear() {
	q.Lock()
	defer q.Unlock()

	q.data = make([]V, 0)
}

// ForEach iterates over the queue and calls the given function for each element from front to back
func (q *Queue[V]) ForEach(fn func(value V, index int) (stop bool)) {
	q.RLock()
	defer q.RUnlock()

	for i, value := range q.data {
		if stop := fn(value, i); stop {
			break
		}
	}
}

// ForEachReverse iterates over the queue and calls the given function for each element from back to front
func (q *Queue[V]) ForEachReverse(fn func(value V, index int) (stop bool)) {
	q.RLock()
	defer q.RUnlock()

	j := 0
	for i := len(q.data) - 1; i >= 0; i-- {
		if stop := fn(q.data[i], j); stop {
			break
		}
		j++
	}
}

// String returns a string representation of the queue
func (q *Queue[V]) String() string {
	bytes, err := q.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// MarshalJSON returns the JSON encoding of the queue
func (q *Queue[V]) MarshalJSON() ([]byte, error) {
	q.RLock()
	defer q.RUnlock()

	return json.Marshal(q.data)
}

// UnmarshalJSON decodes the JSON encoding of the queue
func (q *Queue[V]) UnmarshalJSON(data []byte) error {
	q.Lock()
	defer q.Unlock()

	return json.Unmarshal(data, &q.data)
}

// ToSlice returns the origin queue
func (q *Queue[V]) ToSlice() []V {
	q.RLock()
	defer q.RUnlock()

	return q.data
}
