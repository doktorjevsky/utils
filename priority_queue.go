package utils

import (
	"errors"
)

/*
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
*/

type PriorityQueue[T interface{}] struct {
	comparator func(T, T) int
	contents   []T
}

func NewPriorityQueue[T interface{}](comp func(T, T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		comparator: comp,
		contents:   make([]T, 0)}
}

/*
O(log n)
Assumes: PriorityQueue has been initiated
Inserts the item into a binary heap
*/
func (q *PriorityQueue[T]) Enqueue(item T) {
	pos := len(q.contents)
	q.contents = append(q.contents, item)
	done := false
	parentIndex := getParentIndex(pos)
	for parentIndex >= 0 && !done {
		if q.comparator(q.contents[parentIndex], q.contents[pos]) > 0 {
			parent := q.contents[parentIndex]
			child := q.contents[pos]
			q.contents[parentIndex] = child
			q.contents[pos] = parent
			pos = parentIndex
			parentIndex = getParentIndex(pos)
		} else {
			done = true
		}
	}
}

/*
O(n log n)
Assumes: the priority queue has bee instantiated
Enqueues all items
*/
func (q *PriorityQueue[T]) EnqueueAll(items []T) {
	for _, item := range items {
		q.Enqueue(item)
	}
}

/*
O(log n)
Assumes: the priority queue has been instantiated
Removes the top element and restores the heap invariant
*/
func (q *PriorityQueue[T]) Dequeue() (T, error) {
	var nilVal T
	if len(q.contents) == 0 {
		return nilVal, errors.New(dequeueErrorMsg)
	}
	item := q.contents[0]
	n := len(q.contents)
	pos := 0
	q.contents[pos] = q.contents[n-1]
	q.contents = q.contents[:n-1]
	done := false
	for !done {
		swap := q.getChildSwapIndex(pos)
		if swap < 0 {
			done = true
		} else {
			parent := q.contents[pos]
			child := q.contents[swap]
			q.contents[swap] = parent
			q.contents[pos] = child
			pos = swap
		}
	}
	return item, nil
}

/*
O(n log n)
Assumes: priority queue has been instantiated
Dequeues all items and returns them as a slice. The items will be sorted by the ordering given by pq.comparator
*/
func (q *PriorityQueue[T]) DequeueAll() []T {
	out := make([]T, 0)
	for len(q.contents) > 0 {
		item, _ := q.Dequeue()
		out = append(out, item)
	}
	return out
}

/*
O(1)
Assumes: the priority queue has been instantiated
Returns the top element without removing it from the queue
*/
func (q PriorityQueue[T]) Peek() (T, error) {
	var nilVal T
	if len(q.contents) == 0 {
		return nilVal, errors.New(peekErrorMsg)
	}
	return q.contents[0], nil
}

/*
O(1)
Assumes: the priority queue has been instantiated
Returns true if the queue is empty
*/
func (q PriorityQueue[T]) IsEmpty() bool {
	return len(q.contents) == 0
}

/*
O(1)
Wipes contents of the queue
*/
func (q *PriorityQueue[T]) Clear() {
	q.contents = make([]T, 0)
}

/*
O(n)
Assumes: the priority queue has been instantiated
Returns the inner representation of the queue as a slice
*/
func (q PriorityQueue[T]) ToSlice() []T {
	out := make([]T, 0)
	for _, val := range q.contents {
		out = append(out, val)
	}
	return out
}

/*
O(1)
Assumes: the priority queue has been instantiated
Returns the number of items in the queue
*/
func (q PriorityQueue[T]) Size() int {
	return len(q.contents)
}

// PRIVATE HELPER FUNCTIONS BELOW

// index or -1 if no swap
func (q *PriorityQueue[T]) getChildSwapIndex(parentIndex int) int {
	right := getRightChild(parentIndex)
	left := getLeftChild(parentIndex)
	// no children
	if left > len(q.contents)-1 {
		return -1
		// 1 child
	} else if right > len(q.contents)-1 {
		if q.comparator(q.contents[parentIndex], q.contents[left]) > 0 {
			return left
		} else {
			return -1
		}
		// 2 larger children
	} else if !(q.comparator(q.contents[parentIndex], q.contents[left]) > 0 || q.comparator(q.contents[parentIndex], q.contents[right]) > 0) {
		return -1
		// at least 1 smaller child: is it left?
	} else if q.comparator(q.contents[parentIndex], q.contents[left]) > 0 {
		// left is larger?
		if q.comparator(q.contents[left], q.contents[right]) > 0 {
			// then right is the smallest
			return right
		} else {
			return left
		}
		// right must be smaller than parent and the smallest child
	} else {
		return right
	}

}

func getLeftChild(i int) int {
	return 2*i + 1
}

func getRightChild(i int) int {
	return 2*i + 2
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}
