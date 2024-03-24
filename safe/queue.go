package safe

import "sync"

// Queue ...
type Queue[V any] struct {
	sync.RWMutex
	data []V
}

// NewQueue returns a new safe queue
func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{
		data: make([]V, 0),
	}
}

// Enqueue adds an element to the end of the queue
func (q *Queue[V]) Enqueue(value V) {
	q.Lock()
	defer q.Unlock()

	q.data = append(q.data, value)
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
