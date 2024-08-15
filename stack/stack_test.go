package stack_test

import (
	"testing"

	"github.com/agustfricke/dsa-go/stack"
)

func TestStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		s := stack.NewStack[int]()
		if s.Length != 0 || s.Top != nil {
			t.Errorf("NewStack() did not initialize an empty stack")
		}
	})

	t.Run("Push", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		if s.Length != 1 || s.Top.Value != 1 {
			t.Errorf("Push() did not add element correctly")
		}
		s.Push(2)
		if s.Length != 2 || s.Top.Value != 2 {
			t.Errorf("Push() did not add second element correctly")
		}
	})

	t.Run("Pop", func(t *testing.T) {
		s := stack.NewStack[int]()
		s.Push(1)
		s.Push(2)
		val, err := s.Pop()
		if err != nil || val != 2 || s.Length != 1 {
			t.Errorf("Pop() did not remove element correctly")
		}
		val, err = s.Pop()
		if err != nil || val != 1 || s.Length != 0 {
			t.Errorf("Pop() did not remove last element correctly")
		}
		_, err = s.Pop()
		if err == nil {
			t.Errorf("Pop() should return error on empty stack")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		s := stack.NewStack[int]()
		_, err := s.Peek()
		if err == nil {
			t.Errorf("Peek() should return error on empty stack")
		}
		s.Push(1)
		val, err := s.Peek()
		if err != nil || val != 1 || s.Length != 1 {
			t.Errorf("Peek() did not return correct value")
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		s := stack.NewStack[int]()
		if !s.IsEmpty() {
			t.Errorf("IsEmpty() should return true for new stack")
		}
		s.Push(1)
		if s.IsEmpty() {
			t.Errorf("IsEmpty() should return false after push")
		}
		s.Pop()
		if !s.IsEmpty() {
			t.Errorf("IsEmpty() should return true after popping all elements")
		}
	})

	t.Run("Size", func(t *testing.T) {
		s := stack.NewStack[int]()
		if s.Size() != 0 {
			t.Errorf("Size() should return 0 for new stack")
		}
		s.Push(1)
		if s.Size() != 1 {
			t.Errorf("Size() should return 1 after one push")
		}
		s.Push(2)
		if s.Size() != 2 {
			t.Errorf("Size() should return 2 after two pushes")
		}
		s.Pop()
		if s.Size() != 1 {
			t.Errorf("Size() should return 1 after one pop")
		}
	})
}
