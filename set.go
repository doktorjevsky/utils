package utils

type Set[T comparable] interface {
	Add(item T) bool
	AddAll(item []T)
	Remove(item T) bool
	RemoveAll(item []T)
	Clear()
	Size() int
	ToSlice() []T
}

type HashSet[T comparable] struct {
	items map[T]bool
}

/*
O(1)
Adds item to the set. Returns true if the item wasn't there before
*/
func (s *HashSet[T]) Add(item T) bool {
	res := s.items[item]
	s.items[item] = true
	return res
}

func (s *HashSet[T]) AddAll(items []T) {
	for _, item := range items {
		s.items[item] = true
	}
}

/*
O(1)
Removes item from set. Returns true if the item was there
If the item isn't there, remove is a no-op
*/
func (s *HashSet[T]) Remove(item T) bool {
	if s.items[item] {
		delete(s.items, item)
		return true
	} else {
		return false
	}
}

/*
O(n)
Removes all items in the argument list
*/
func (s *HashSet[T]) RemoveAll(items []T) {
	for _, i := range items {
		delete(s.items, i)
	}
}

/*
O(1)
Clears the set
*/
func (s *HashSet[T]) Clear() {
	s.items = make(map[T]bool)
}

func (s HashSet[T]) Size() int {
	return len(s.items)
}

func (s HashSet[T]) ToSlice() []T {
	out := make([]T, 0)
	for k, _ := range s.items {
		out = append(out, k)
	}
	return out
}
