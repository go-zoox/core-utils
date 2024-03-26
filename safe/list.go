package safe

import (
	"encoding/json"
	"reflect"
	"sync"
)

// List ...
type List[V any] struct {
	sync.RWMutex
	data []V
	//
	cfg *ListConfig
}

// ListConfig ...
type ListConfig struct {
	Capacity int
}

// ListOption ...
type ListOption func(*ListConfig)

// NewList returns a new safe list
//
// @TODO generic type comparable interface cannot implement by struct，so we use any type for generic type
func NewList[V any](opts ...ListOption) *List[V] {
	cfg := &ListConfig{
		Capacity: 0,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return &List[V]{
		cfg:  cfg,
		data: make([]V, 0),
	}
}

// Push adds an element to the end of the list
func (l *List[V]) Push(value V) {
	l.Lock()
	defer l.Unlock()

	l.data = append(l.data, value)

	// check capacity when push
	if l.cfg.Capacity > 0 && len(l.data) > l.cfg.Capacity {
		l.data = l.data[len(l.data)-l.cfg.Capacity:]
	}
}

// LPush adds an element to the beginning of the list
func (l *List[V]) LPush(value V) {
	l.Lock()
	defer l.Unlock()

	l.data = append([]V{value}, l.data...)

	// check capacity when push
	if l.cfg.Capacity > 0 && len(l.data) > l.cfg.Capacity {
		l.data = l.data[:l.cfg.Capacity]
	}
}

// Pop removes and returns the last element of the list
func (l *List[V]) Pop() V {
	l.Lock()
	defer l.Unlock()

	if len(l.data) == 0 {
		var v V
		return v
	}

	value := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-1]
	return value
}

// LPop removes and returns the first element of the list
func (l *List[V]) LPop() V {
	l.Lock()
	defer l.Unlock()

	if len(l.data) == 0 {
		var v V
		return v
	}

	value := l.data[0]
	l.data = l.data[1:]
	return value
}

// Size returns the number of elements in the list
func (l *List[V]) Size() int {
	l.RLock()
	defer l.RUnlock()

	return len(l.data)
}

// Length is an alias of Size
func (l *List[V]) Length() int {
	return l.Size()
}

// Clear removes all elements from the list
func (l *List[V]) Clear() {
	l.Lock()
	defer l.Unlock()

	l.data = make([]V, 0)
}

// Get returns the element at the given index
func (l *List[V]) Get(index int) V {
	l.RLock()
	defer l.RUnlock()

	if index < 0 || index >= len(l.data) {
		var v V
		return v
	}

	return l.data[index]
}

// ForEach iterates over the list and calls the given function for each element
func (l *List[V]) ForEach(f func(V) (stop bool)) {
	l.RLock()
	defer l.RUnlock()

	for _, value := range l.data {
		if stop := f(value); stop {
			break
		}
	}
}

// Filter returns a new list with all elements that satisfy the given function
func (l *List[V]) Filter(f func(V) bool) *List[V] {
	l.RLock()
	defer l.RUnlock()

	newList := NewList[V]()
	for _, value := range l.data {
		if f(value) {
			newList.Push(value)
		}
	}

	return newList
}

// Reduce returns the result of reducing the list to a single value
func (l *List[V]) Reduce(f func(V, V) V) V {
	l.RLock()
	defer l.RUnlock()

	if len(l.data) == 0 {
		var v V
		return v
	}

	result := l.data[0]
	for _, value := range l.data[1:] {
		result = f(result, value)
	}

	return result
}

// Map returns a new list with the result of calling the given function on each element
func (l *List[V]) Map(f func(V) V) *List[V] {
	l.RLock()
	defer l.RUnlock()

	newList := NewList[V]()
	for _, value := range l.data {
		newList.Push(f(value))
	}

	return newList
}

// IndexOf returns the index of the given element
func (l *List[V]) IndexOf(value V) int {
	l.RLock()
	defer l.RUnlock()

	for i, v := range l.data {
		// @TODO
		// if v == value {
		// 	return i
		// }
		if reflect.DeepEqual(v, value) {
			return i
		}
	}

	return -1
}

// Find returns the first element that satisfies the given function
func (l *List[V]) Find(f func(V) bool) V {
	l.RLock()
	defer l.RUnlock()

	for _, value := range l.data {
		if f(value) {
			return value
		}
	}

	var v V
	return v
}

// Contains returns true if the list contains the given element
func (l *List[V]) Contains(value V) bool {
	l.RLock()
	defer l.RUnlock()

	return l.IndexOf(value) != -1
}

// First returns the first element of the list
func (l *List[V]) First() V {
	l.RLock()
	defer l.RUnlock()

	if len(l.data) == 0 {
		var v V
		return v
	}

	return l.data[0]
}

// Last returns the last element of the list
func (l *List[V]) Last() V {
	l.RLock()
	defer l.RUnlock()

	if len(l.data) == 0 {
		var v V
		return v
	}

	return l.data[len(l.data)-1]
}

// Unshift adds an element to the beginning of the list
func (l *List[V]) Unshift(value V) {
	l.Lock()
	defer l.Unlock()

	l.data = append([]V{value}, l.data...)

	// check capacity when push
	if l.cfg.Capacity > 0 && len(l.data) > l.cfg.Capacity {
		l.data = l.data[:l.cfg.Capacity]
	}
}

// Reverse reverses the list
func (l *List[V]) Reverse() *List[V] {
	l.Lock()
	defer l.Unlock()

	newList := NewList[V]()
	for i := len(l.data) - 1; i >= 0; i-- {
		newList.Push(l.data[i])
	}
	return newList
}

// Splice removes the given number of elements from the given index
func (l *List[V]) Splice(index int, count int) *List[V] {
	l.Lock()
	defer l.Unlock()

	if index < 0 || index >= len(l.data) {
		return nil
	}

	if count < 0 || count > len(l.data)-index {
		count = len(l.data) - index
	}

	newList := NewList[V]()
	for i := 0; i < count; i++ {
		newList.Push(l.data[index+i])
	}

	l.data = append(l.data[:index], l.data[index+count:]...)
	return newList
}

// Slice returns a new list with a copy of a given number of elements from the given index
func (l *List[V]) Slice(start int, end int) *List[V] {
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

	newList := NewList[V]()
	for i := start; i < end; i++ {
		newList.Push(l.data[i])
	}

	return newList
}

// Swap swaps the elements at the given indices
func (l *List[V]) Swap(index1 int, index2 int) {
	l.Lock()
	defer l.Unlock()

	if index1 < 0 || index1 >= len(l.data) || index2 < 0 || index2 >= len(l.data) {
		return
	}

	l.data[index1], l.data[index2] = l.data[index2], l.data[index1]
}

// Iterator returns a channel that will yield successive elements of the list
func (l *List[V]) Iterator() []V {
	return l.ToSlice()
}

// ToSlice returns the origin map
func (l *List[V]) ToSlice() []V {
	l.RLock()
	defer l.RUnlock()

	return l.data
}

// String returns the string representation of the list
func (l *List[V]) String() string {
	bytes, err := l.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// MarshalJSON returns the JSON encoding of the list
func (l *List[V]) MarshalJSON() ([]byte, error) {
	l.RLock()
	defer l.RUnlock()

	return json.Marshal(l.data)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in the list
func (l *List[V]) UnmarshalJSON(data []byte) error {
	l.Lock()
	defer l.Unlock()

	return json.Unmarshal(data, &l.data)
}

// Unshift adds an element to the beginning of the list
//
//	like JavaScript Array.unshift
func (l *List[V]) Unshifts(values V) {
	l.LPush(values)
}

// Shift removes and returns the first element of the list
//
//	like JavaScript Array.shift
func (l *List[V]) Shift() V {
	return l.LPop()
}
