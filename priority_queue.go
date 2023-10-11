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

// parent of i at (i-1) / 2
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

func getParentIndex(i int) int {
	return (i - 1) / 2
}

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
