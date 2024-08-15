package llqueue

import (
	"errors"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{Value: value}
	if q.IsEmpty() {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
	q.Length++
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	q.Length--
	head := q.Head
	q.Head = q.Head.Next
	head.Next = nil
	if q.IsEmpty() {
		q.Tail = nil
	}
	return head.Value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	return q.Head.Value, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Head == nil
}
