package utils

import "errors"

const emptyStackError = "Stack is empty"

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

type SliceStack[T interface{}] struct {
	items []T
}

func NewSliceStack[T interface{}]() *SliceStack[T] {
	return &SliceStack[T]{items: make([]T, 0)}
}

func (s *SliceStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *SliceStack[T]) PushAll(items []T) {
	for _, item := range items {
		s.items = append(s.items, item)
	}
}

func (s *SliceStack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var nilVal T
		return nilVal, errors.New(emptyStackError)
	}
	n := len(s.items)
	item := s.items[n-1]

	s.items = s.items[:n-1]
	return item, nil
}

func (s *SliceStack[T]) PopAll() []T {
	out := make([]T, len(s.items))
	for i, val := range s.items {
		out[len(s.items)-i-1] = val
	}
	s.items = make([]T, 0)
	return out
}

func (s SliceStack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var nilVal T
		return nilVal, errors.New(emptyStackError)
	}
	return s.items[len(s.items)-1], nil
}

func (s SliceStack[T]) Size() int {
	return len(s.items)
}

func (s SliceStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s SliceStack[T]) ToSlice() []T {
	out := make([]T, len(s.items))
	for i, v := range s.items {
		out[i] = v
	}
	return out
}

func (s *SliceStack[T]) Clear() {
	s.items = make([]T, 0)
}
