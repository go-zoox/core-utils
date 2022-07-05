package safe

import (
	"sync"
)

type List struct {
	sync.RWMutex
	data []interface{}
}

// NewList returns a new safe list
func NewList() *List {
	return &List{
		data: make([]interface{}, 0),
	}
}

// Push adds an element to the end of the list
func (l *List) Push(value interface{}) {
	l.Lock()
	defer l.Unlock()

	l.data = append(l.data, value)
}

// Pop removes and returns the last element of the list
func (l *List) Pop() interface{} {
	l.Lock()
	defer l.Unlock()

	if len(l.data) == 0 {
		return nil
	}

	value := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-1]
	return value
}

// Size returns the number of elements in the list
func (l *List) Size() int {
	l.RLock()
	defer l.RUnlock()

	return len(l.data)
}

// Length is an alias of Size
func (l *List) Length() int {
	return l.Size()
}

// Clear removes all elements from the list
func (l *List) Clear() {
	l.Lock()
	defer l.Unlock()

	l.data = make([]interface{}, 0)
}

// Get returns the element at the given index
func (l *List) Get(index int) interface{} {
	l.RLock()
	defer l.RUnlock()

	if index < 0 || index >= len(l.data) {
		return nil
	}

	return l.data[index]
}

// ForEach iterates over the list and calls the given function for each element
func (l *List) ForEach(f func(interface{})) {
	l.RLock()
	defer l.RUnlock()

	for _, value := range l.data {
		f(value)
	}
}

// Filter returns a new list with all elements that satisfy the given function
func (l *List) Filter(f func(interface{}) bool) *List {
	l.RLock()
	defer l.RUnlock()

	newList := NewList()
	for _, value := range l.data {
		if f(value) {
			newList.Push(value)
		}
	}

	return newList
}

// Reduce returns the result of reducing the list to a single value
func (l *List) Reduce(f func(interface{}, interface{}) interface{}) interface{} {
	l.RLock()
	defer l.RUnlock()

	if len(l.data) == 0 {
		return nil
	}

	result := l.data[0]
	for _, value := range l.data[1:] {
		result = f(result, value)
	}

	return result
}

// Map returns a new list with the result of calling the given function on each element
func (l *List) Map(f func(interface{}) interface{}) *List {
	l.RLock()
	defer l.RUnlock()

	newList := NewList()
	for _, value := range l.data {
		newList.Push(f(value))
	}

	return newList
}

// IndexOf returns the index of the given element
func (l *List) IndexOf(value interface{}) int {
	l.RLock()
	defer l.RUnlock()

	for i, v := range l.data {
		if v == value {
			return i
		}
	}

	return -1
}

// Find returns the first element that satisfies the given function
func (l *List) Find(f func(interface{}) bool) interface{} {
	l.RLock()
	defer l.RUnlock()

	for _, value := range l.data {
		if f(value) {
			return value
		}
	}

	return nil
}

// Contains returns true if the list contains the given element
func (l *List) Contains(value interface{}) bool {
	l.RLock()
	defer l.RUnlock()

	return l.IndexOf(value) != -1
}

// First returns the first element of the list
func (l *List) First() interface{} {
	l.RLock()
	defer l.RUnlock()

	if len(l.data) == 0 {
		return nil
	}

	return l.data[0]
}

// Last returns the last element of the list
func (l *List) Last() interface{} {
	l.RLock()
	defer l.RUnlock()

	if len(l.data) == 0 {
		return nil
	}

	return l.data[len(l.data)-1]
}

// Unshift adds an element to the beginning of the list
func (l *List) Unshift(value interface{}) {
	l.Lock()
	defer l.Unlock()

	l.data = append([]interface{}{value}, l.data...)
}

// Shift removes and returns the first element of the list
func (l *List) Shift() interface{} {
	l.Lock()
	defer l.Unlock()

	if len(l.data) == 0 {
		return nil
	}

	value := l.data[0]
	l.data = l.data[1:]
	return value
}

// Reverse reverses the list
func (l *List) Reverse() *List {
	l.Lock()
	defer l.Unlock()

	newList := NewList()
	for i := len(l.data) - 1; i >= 0; i-- {
		newList.Push(l.data[i])
	}
	return newList
}

// Splice removes the given number of elements from the given index
func (l *List) Splice(index int, count int) *List {
	l.Lock()
	defer l.Unlock()

	if index < 0 || index >= len(l.data) {
		return nil
	}

	if count < 0 || count > len(l.data)-index {
		count = len(l.data) - index
	}

	newList := NewList()
	for i := 0; i < count; i++ {
		newList.Push(l.data[index+i])
	}

	l.data = append(l.data[:index], l.data[index+count:]...)
	return newList
}

// Slice returns a new list with a copy of a given number of elements from the given index
func (l *List) Slice(start int, end int) *List {
	l.RLock()
	defer l.RUnlock()

	if start < 0 || start >= len(l.data) {
		return nil
	}

	if end < 0 || end >= len(l.data) {
		end = len(l.data)
	}

	if start > end {
		start, end = end, start
	}

	newList := NewList()
	for i := start; i < end; i++ {
		newList.Push(l.data[i])
	}

	return newList
}

// Swap swaps the elements at the given indices
func (l *List) Swap(index1 int, index2 int) {
	l.Lock()
	defer l.Unlock()

	if index1 < 0 || index1 >= len(l.data) || index2 < 0 || index2 >= len(l.data) {
		return
	}

	l.data[index1], l.data[index2] = l.data[index2], l.data[index1]
}

// Iterator returns a channel that will yield successive elements of the list
func (l *List) Iterator() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		l.RLock()
		defer l.RUnlock()

		for _, value := range l.data {
			c <- value
		}
		close(c)
	}()
	return c
}
