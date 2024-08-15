package arrayqueue

import (
  "errors"
)

type Queue struct {
  items [5]int
  front int
  rear  int
  size  int
}

func NewQueue() *Queue {
  return &Queue{
    front: 0,
    rear:  -1,
    size:  0,
  }
}

func (q *Queue) Enqueue(item int) error {
  if q.IsFull() {
    return errors.New("Cola llena")
  }
  q.rear = (q.rear + 1) % 5
  q.items[q.rear] = item
  q.size++
  return nil
}

func (q *Queue) Dequeue() (int, error) {
  if q.IsEmpty() {
    return 0, errors.New("Cola vacía")
  }
  item := q.items[q.front]
  q.front = (q.front + 1) % 5
  q.size--
  return item, nil
}

func (q *Queue) Front() (int, error) {
  if q.IsEmpty() {
    return 0, errors.New("Cola vacía")
  }
  return q.items[q.front], nil
}

func (q *Queue) IsEmpty() bool {
  return q.size == 0
}

func (q *Queue) IsFull() bool {
  return q.size == 5
}

func (q *Queue) Size() int {
  return q.size
}
