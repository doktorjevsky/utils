package utils

import (
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

	for i := 0; i < 99/2; i++ {
		if !(pq.contents[i] <= pq.contents[2*i+1] && pq.contents[i] <= pq.contents[2*i+2]) {
			t.Errorf("Heap invariant is broken. Parent: %d | Left child: %d | Right child: %d", pq.contents[i], pq.contents[2*i+1], pq.contents[2*i+2])
		}
	}
}
