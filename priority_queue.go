package utils

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
