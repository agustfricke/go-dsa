package stack

import "errors"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	Top    *Node[T]
	Length int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{Value: value, Next: s.Top}
	s.Top = newNode
	s.Length++
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	value := s.Top.Value
	s.Top = s.Top.Next
	s.Length--
	return value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.Top.Value, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length == 0
}

func (s *Stack[T]) Size() int {
	return s.Length
}
