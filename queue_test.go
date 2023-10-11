package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestFifoQueue_Enqueue(t *testing.T) {
	// Create a new FifoQueue
	q := NewFifoQueue[int]()

	// Enqueue a few items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue items and verify their order
	item, _ := q.Dequeue()
	if item != 1 {
		t.Errorf("Dequeued item is not as expected. Got %v, expected 1", item)
	}

	item, _ = q.Dequeue()
	if item != 2 {
		t.Errorf("Dequeued item is not as expected. Got %v, expected 2", item)
	}

	item, _ = q.Dequeue()
	if item != 3 {
		t.Errorf("Dequeued item is not as expected. Got %v, expected 3", item)
	}

	item, err := q.Dequeue()
	if err == nil {
		t.Fail()
	}
}

func TestFifoQueue_EnqueueAll(t *testing.T) {
	// Create a new FifoQueue
	q := NewFifoQueue[int]()

	// Enqueue a few items using EnqueueAll
	items := []int{1, 2, 3}
	q.EnqueueAll(items)

	// Dequeue items and verify their order
	item, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue returned an error: %v", err)
	}
	if item != 1 {
		t.Errorf("Dequeued item is not as expected. Got %v, expected 1", item)
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue returned an error: %v", err)
	}
	if item != 2 {
		t.Errorf("Dequeued item is not as expected. Got %v, expected 2", item)
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue returned an error: %v", err)
	}
	if item != 3 {
		t.Errorf("Dequeued item is not as expected. Got %v, expected 3", item)
	}
}

func TestFifoQueue_DequeueAll(t *testing.T) {
	// Create a new FifoQueue
	q := NewFifoQueue[int]()

	// Enqueue a few items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue all items and verify their order
	items := q.DequeueAll()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(items, expected) {
		t.Errorf("Dequeued items are not as expected. Got %v, expected %v", items, expected)
	}

	// Verify that the queue is empty after dequeueing all items
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after DequeueAll")
	}
}

func TestFifoQueue_Peek(t *testing.T) {
	// Create a new FifoQueue
	q := NewFifoQueue[int]()

	// Attempt to Peek when the queue is empty
	item, err := q.Peek()
	expectedError := errors.New(peekErrorMsg)
	if !reflect.DeepEqual(err.Error(), expectedError.Error()) {
		t.Errorf("Peek should return an error when the queue is empty. Got error: %v, expected error: %v", err, expectedError)
	}

	// Enqueue some items
	q.Enqueue(1)
	q.Enqueue(2)

	// Peek at the front item without removing it
	item, err = q.Peek()
	if err != nil {
		t.Errorf("Peek should not return an error when the queue is not empty. Got error: %v", err)
	}
	if item != 1 {
		t.Errorf("Peek should return the first enqueued item. Got %v, expected 1", item)
	}

	// Verify that the queue is not empty after peeking
	if q.IsEmpty() {
		t.Errorf("Queue should not be empty after Peek")
	}
}

func TestFifoQueue_ToSlice(t *testing.T) {
	// Create a new FifoQueue
	q := NewFifoQueue[int]()

	// Enqueue some items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Get the queue as a slice
	slice := q.ToSlice()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("ToSlice should return the queue as a slice. Got %v, expected %v", slice, expected)
	}
}
