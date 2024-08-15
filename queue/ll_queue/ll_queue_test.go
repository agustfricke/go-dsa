package llqueue_test

import (
	"testing"

	llqueue "github.com/agustfricke/dsa-go/queue/ll_queue"
)

func TestQueue(t *testing.T) {
	t.Run("NewQueue", func(t *testing.T) {
		q := llqueue.NewQueue[int]()
		if q.Length != 0 || q.Head != nil || q.Tail != nil {
			t.Errorf("NewQueue() did not initialize an empty queue")
		}
	})

	t.Run("Enqueue", func(t *testing.T) {
		q := llqueue.NewQueue[int]()
		q.Enqueue(1)
		if q.Length != 1 || q.Head.Value != 1 || q.Tail.Value != 1 {
			t.Errorf("Enqueue() did not add element correctly")
		}
		q.Enqueue(2)
		if q.Length != 2 || q.Head.Value != 1 || q.Tail.Value != 2 {
			t.Errorf("Enqueue() did not add second element correctly")
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := llqueue.NewQueue[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		val, err := q.Dequeue()
		if err != nil || val != 1 || q.Length != 1 {
			t.Errorf("Dequeue() did not remove element correctly")
		}
		val, err = q.Dequeue()
		if err != nil || val != 2 || q.Length != 0 {
			t.Errorf("Dequeue() did not remove last element correctly")
		}
		_, err = q.Dequeue()
		if err == nil {
			t.Errorf("Dequeue() should return error on empty queue")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		q := llqueue.NewQueue[int]()
		_, err := q.Peek()
		if err == nil {
			t.Errorf("Peek() should return error on empty queue")
		}
		q.Enqueue(1)
		val, err := q.Peek()
		if err != nil || val != 1 || q.Length != 1 {
			t.Errorf("Peek() did not return correct value")
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		q := llqueue.NewQueue[int]()
		if !q.IsEmpty() {
			t.Errorf("IsEmpty() should return true for new queue")
		}
		q.Enqueue(1)
		if q.IsEmpty() {
			t.Errorf("IsEmpty() should return false after enqueue")
		}
		q.Dequeue()
		if !q.IsEmpty() {
			t.Errorf("IsEmpty() should return true after dequeuing all elements")
		}
	})
}
