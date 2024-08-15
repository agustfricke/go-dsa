package arrayqueue_test

import (
  "testing"

  arrayqueue "github.com/agustfricke/dsa-go/queue/array_queue"
)

func TestNewQueue(t *testing.T) {
  q := arrayqueue.NewQueue()
  if q.Size() != 0 {
    t.Errorf("Expected empty queue, got size %d", q.Size())
  }
}

func TestEnqueue(t *testing.T) {
  q := arrayqueue.NewQueue()
  err := q.Enqueue(1)
  if err != nil {
    t.Errorf("Unexpected error: %v", err)
  }
  if q.Size() != 1 {
    t.Errorf("Expected size 1, got %d", q.Size())
  }
}

func TestEnqueueFull(t *testing.T) {
  q := arrayqueue.NewQueue()
  for i := 0; i < 5; i++ {
    q.Enqueue(i)
  }
  err := q.Enqueue(5)
  if err == nil {
    t.Error("Expected error, got nil")
  }
}

func TestDequeue(t *testing.T) {
  q := arrayqueue.NewQueue()
  q.Enqueue(1)
  item, err := q.Dequeue()
  if err != nil {
    t.Errorf("Unexpected error: %v", err)
  }
  if item != 1 {
    t.Errorf("Expected 1, got %d", item)
  }
  if q.Size() != 0 {
    t.Errorf("Expected size 0, got %d", q.Size())
  }
}

func TestDequeueEmpty(t *testing.T) {
  q := arrayqueue.NewQueue()
  _, err := q.Dequeue()
  if err == nil {
    t.Error("Expected error, got nil")
  }
}

func TestFront(t *testing.T) {
  q := arrayqueue.NewQueue()
  q.Enqueue(1)
  item, err := q.Front()
  if err != nil {
    t.Errorf("Unexpected error: %v", err)
  }
  if item != 1 {
    t.Errorf("Expected 1, got %d", item)
  }
  if q.Size() != 1 {
    t.Errorf("Expected size 1, got %d", q.Size())
  }
}

func TestFrontEmpty(t *testing.T) {
  q := arrayqueue.NewQueue()
  _, err := q.Front()
  if err == nil {
    t.Error("Expected error, got nil")
  }
}

func TestIsEmpty(t *testing.T) {
  q := arrayqueue.NewQueue()
  if !q.IsEmpty() {
    t.Error("Expected queue to be empty")
  }
  q.Enqueue(1)
  if q.IsEmpty() {
    t.Error("Expected queue to be not empty")
  }
}

func TestIsFull(t *testing.T) {
  q := arrayqueue.NewQueue()
  if q.IsFull() {
    t.Error("Expected queue to be not full")
  }
  for i := 0; i < 5; i++ {
    q.Enqueue(i)
  }
  if !q.IsFull() {
    t.Error("Expected queue to be full")
  }
}

func TestSize(t *testing.T) {
  q := arrayqueue.NewQueue()
  if q.Size() != 0 {
    t.Errorf("Expected size 0, got %d", q.Size())
  }
  q.Enqueue(1)
  if q.Size() != 1 {
    t.Errorf("Expected size 1, got %d", q.Size())
  }
}
