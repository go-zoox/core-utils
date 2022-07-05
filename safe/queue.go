package safe

import "sync"

type Queue struct {
	sync.RWMutex
	data []interface{}
}

// NewQueue returns a new safe queue
func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
	}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.Lock()
	defer q.Unlock()

	q.data = append(q.data, value)
}

// Dequeue removes and returns the first element of the queue
func (q *Queue) Dequeue() interface{} {
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
func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.data)
}

// Front returns the first element of the queue
func (q *Queue) Front() interface{} {
	q.RLock()
	defer q.RUnlock()

	if len(q.data) == 0 {
		return nil
	}

	return q.data[0]
}

// Back returns the last element of the queue
func (q *Queue) Back() interface{} {
	q.RLock()
	defer q.RUnlock()

	if len(q.data) == 0 {
		return nil
	}

	return q.data[len(q.data)-1]
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.data) == 0
}

// Clear removes all elements from the queue
func (q *Queue) Clear() {
	q.Lock()
	defer q.Unlock()

	q.data = make([]interface{}, 0)
}
