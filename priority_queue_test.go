package utils

import (
	"errors"
	"math/rand"
	"testing"
)

var abs = func(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}
var cmp = func(a int, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

var cmp2 = func(a int, b int) int { return cmp(b, a) }

func TestEnqueue(t *testing.T) {
	pq := NewPriorityQueue[int](cmp)
	pq.Enqueue(5)
	pq.Enqueue(2)
	pq.Enqueue(10)

	parent := pq.contents[0]
	child1 := pq.contents[1]
	child2 := pq.contents[2]
	if !(parent <= child1 && parent <= child2) {
		t.Errorf("Expected [2 5 10] or [2 10 5] but was %v", pq.contents)
	}
}

func TestEnqueue2(t *testing.T) {
	pq := NewPriorityQueue[int](cmp2)
	pq.Enqueue(5)
	pq.Enqueue(2)
	pq.Enqueue(10)

	parent := pq.contents[0]
	child1 := pq.contents[1]
	child2 := pq.contents[2]
	if !(parent >= child1 && parent >= child2) {
		t.Errorf("Expected [10 5 2] or [10 2 5] but was %v", pq.contents)
	}

}

func TestEnqueue3(t *testing.T) {
	pq := NewPriorityQueue[int](cmp)
	for i := 0; i < 100; i++ {
		x := rand.Int()
		pq.Enqueue(x)
	}

	err, is := pqInvariant[int](*pq)
	if err != nil {
		t.Errorf("%s. Parent: %d | Left: %d | Right: %d", err.Error(), pq.contents[is[0]], pq.contents[is[1]], pq.contents[is[2]])
	}
}

func TestDequeue(t *testing.T) {
	pq := NewPriorityQueue[int](cmp)
	pq.Enqueue(5)
	pq.Enqueue(10)
	pq.Enqueue(1)
	pq.Enqueue(23)
	pq.Enqueue(-10)

	expected := []int{-10, 1, 5, 10, 23}
	for i := 0; i < len(expected); i++ {
		actual, _ := pq.Dequeue()
		if actual != expected[i] {
			t.Errorf("Expected: %d, but Actual was: %d | HEAP: %v", expected[i], actual, pq.contents)
		}
		if len(pq.contents) != 4-i {
			t.Errorf("Expected queue length: %d, but was given: %d for heap: %v", 4-i, len(pq.contents), pq.contents)
		}
	}
}

func TestDequeue2(t *testing.T) {
	pq := NewPriorityQueue[int](cmp)
	for i := 0; i < 100; i++ {
		for j := 0; j < 1000; j++ {
			pq.Enqueue(rand.Int())
		}
		err, is := pqInvariant[int](*pq)
		if err != nil {
			t.Errorf("%s. Parent: %d | Right: %d | Left: %d", err.Error(), pq.contents[is[0]], pq.contents[is[1]], pq.contents[is[2]])
		}

	}

}

func pqInvariant[T interface{}](pq PriorityQueue[T]) (error, []int) {
	for i := 0; i < (len(pq.contents)-1)/2; i++ {
		if !(pq.comparator(pq.contents[i], pq.contents[2*i+1]) <= 0 && pq.comparator(pq.contents[i], pq.contents[2*i+2]) <= 0) {
			return errors.New("Heap invariant broken"), []int{i, 2*i + 1, 2*i + 2}
		}
	}
	return nil, []int{}
}
