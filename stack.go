package utils

type Stack[T interface{}] interface {
	Push(item T)
	PushAll(items []T)
	Pop() (T, error)
	PopAll() []T
	Peek() (T, error)
	Size() int
	IsEmpty() bool
	ToSlice() []T
	Clear()
}
