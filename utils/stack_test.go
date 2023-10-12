package utils

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	var s Stack[int] = NewSliceStack[int]()
	test := []int{1, 2, 3, 4, 5}
	s.PushAll(test)
	if s.Size() != 5 {
		t.Errorf("Expected size: %d, but actual was: %d for input %v", 5, s.Size(), test)
	}
	result := s.ToSlice()
	for i, v := range test {
		if result[i] != v {
			t.Errorf("Expected: %v but got Actual: %v", test, result)
		}
	}
	result = s.PopAll()
	for i, v := range test {
		if result[4-i] != v {
			t.Fail()
		}
	}
}

func TestPushAndPop2(t *testing.T) {
	var s Stack[int] = NewSliceStack[int]()
	s.Push(1)
	res, _ := s.Peek()
	if res != 1 {
		t.Errorf("Expected: %d | Actual: %d", 1, res)
	}
	s.Push(2)
	res, _ = s.Peek()
	if res != 2 {
		t.Errorf("Expected: %d | Actual: %d", 2, res)
	}
	s.Pop()
	res, _ = s.Peek()
	if res != 1 {
		t.Errorf("Expected: %d | Actual: %d", 1, res)
	}
}

func TestClear(t *testing.T) {
	var s Stack[int] = NewSliceStack[int]()
	s.PushAll([]int{1, 2, 3, 4})
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Expected stackto be empty but has size: %d", s.Size())
	}
}
