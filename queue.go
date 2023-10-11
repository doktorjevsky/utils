package utils

import "errors"

const dequeueErrorMsg string = "Cannot dequeue an empty queue"
const peekErrorMsg string = "Cannot peek in an empty queue"

/*
 A generic queue interface
*/
type Queue[T interface{}] interface {
	Enqueue(item T)
	EnqueueAll(item []T)
	Dequeue() (T, error)
	DequeueAll() []T
	Peek() (T, error)
	IsEmpty() bool
	Size() int
	Clear()
	ToSlice() []T
}

/*
 A FifoQueue
*/
type FifoQueue[T interface{}] struct {
	contents []T
}

/*
O(1)
Instantiates a new FifoQueue
*/
func NewFifoQueue[T interface{}]() FifoQueue[T] {
	return FifoQueue[T]{contents: make([]T, 0)}
}

/*
O(1)
Assumes: the queue has been instantiated
Places the item at the back of the queue
*/
func (q *FifoQueue[T]) Enqueue(item T) {
	q.contents = append(q.contents, item)
}

/*
O(n)
Assumes: the queue has been instantiated
Iterates over the items and places each item at the back of the queue
*/
func (q *FifoQueue[T]) EnqueueAll(items []T) {
	for _, item := range items {
		q.contents = append(q.contents, item)
	}
}

/*
O(1)
Assumes: the queue has been instantiated
Removes the item at the front of the queue and returns it
Returns error if the queue is empty
*/
func (q *FifoQueue[T]) Dequeue() (T, error) {
	n := len(q.contents)
	var nilVal T
	if n == 0 {
		return nilVal, errors.New(dequeueErrorMsg)
	}
	item := q.contents[0]
	q.contents = q.contents[1:]
	return item, nil
}

/*
O(n)
Assumes: the queue has been instantiated
Removes all items in the queue and returns them as a slice in the order they were placed in the queue
*/
func (q *FifoQueue[T]) DequeueAll() []T {
	out := make([]T, 0)
	for _, val := range q.contents {
		out = append(out, val)
	}
	q.contents = make([]T, 0)
	return out
}

/*
O(1)
Assumes: the queue has been instantiated
Returns the item at the front of the queue without removing it
Returns error if the queue is empty
*/
func (q FifoQueue[T]) Peek() (T, error) {
	var nilVal T
	if len(q.contents) == 0 {
		return nilVal, errors.New(peekErrorMsg)
	}
	return q.contents[0], nil
}

/*
O(1)
Assumes: the queue has been instantiated
Returns true if there are no items in the queue
*/
func (q FifoQueue[T]) IsEmpty() bool {
	return len(q.contents) == 0
}

/*
O(1)
Assumes: the queue has been instantiated
Returns the number of items in the queue
*/
func (q FifoQueue[T]) Size() int {
	return len(q.contents)
}

/*
Replaces the internal queue with an empty queue
*/
func (q FifoQueue[T]) Clear() {
	q.contents = make([]T, 0)
}

/*
Assumes: the queue has been instantiated
Returns the queue as a slice
*/
func (q FifoQueue[T]) ToSlice() []T {
	out := make([]T, 0)
	for _, val := range q.contents {
		out = append(out, val)
	}
	return out
}
