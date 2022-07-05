package safe

import "sync"

type safeQueue struct {
	sync.RWMutex
	data []interface{}
}

// NewQueue returns a new safe queue
func NewQueue() *safeQueue {
	return &safeQueue{
		data: make([]interface{}, 0),
	}
}

// Enqueue adds an element to the end of the queue
func (q *safeQueue) Enqueue(value interface{}) {
	q.Lock()
	defer q.Unlock()

	q.data = append(q.data, value)
}

// Dequeue removes and returns the first element of the queue
func (q *safeQueue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.data) == 0 {
		return nil
	}

	value := q.data[0]
	q.data = q.data[1:]
	return value
}

// Size returns the number of elements in the queue
func (q *safeQueue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.data)
}

// Front returns the first element of the queue
func (q *safeQueue) Front() interface{} {
	q.RLock()
	defer q.RUnlock()

	if len(q.data) == 0 {
		return nil
	}

	return q.data[0]
}

// Back returns the last element of the queue
func (q *safeQueue) Back() interface{} {
	q.RLock()
	defer q.RUnlock()

	if len(q.data) == 0 {
		return nil
	}

	return q.data[len(q.data)-1]
}

// IsEmpty returns true if the queue is empty
func (q *safeQueue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.data) == 0
}

// Clear removes all elements from the queue
func (q *safeQueue) Clear() {
	q.Lock()
	defer q.Unlock()

	q.data = make([]interface{}, 0)
}
