package doublylinkedlist_test

import (
	"testing"

	doublylinkedlist "github.com/agustfricke/dsa-go/doubly_linked_list"
)

func TestDeleteTail(t *testing.T) {
	t.Run("delete tail empty dll", func(t *testing.T) {
		dll := doublylinkedlist.DoublyLinkedList{}
		err := dll.DeleteTail()
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("delete tail one node dll", func(t *testing.T) {
		dll := doublylinkedlist.DoublyLinkedList{}
		dll.InsertAtEnd(73)
		err := dll.DeleteTail()
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if dll.Head != nil {
			t.Fatalf("expected head to be nil but got: %v", dll.Head.Data)
		}
		if dll.Tail != nil {
			t.Fatalf("expected tail to be nil but got %v", dll.Tail.Data)
		}
	})

	t.Run("delete tail with multiples nodes dll", func(t *testing.T) {
		dll := doublylinkedlist.DoublyLinkedList{}
		dll.InsertAtEnd(73)
		dll.InsertAtEnd(69)
		dll.InsertAtEnd(420)
		dll.InsertAtEnd(777)
		err := dll.DeleteTail()
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if dll.Head.Data != 73 {
			t.Fatalf("expected head to be 73 but got %v", dll.Head.Data)
		}
		if dll.Head.Next.Data != 69 {
			t.Fatalf("expected head.next to be 69 but got %v", dll.Head.Next.Data)
		}
		if dll.Tail.Data != 420 {
			t.Fatalf("expected tail to be 420 but got %v", dll.Tail.Data)
		}
		if dll.Tail.Prev.Data != 69 {
			t.Fatalf("expected tail.prev to be 69 but got %v", dll.Tail.Prev.Data)
		}
		if dll.Head.Prev != nil {
			t.Fatalf("expected head.prev to be nil but got %v", dll.Head.Prev.Data)
		}
		if dll.Tail.Next != nil {
			t.Fatalf("expected tail.next to be nil but got %v", dll.Tail.Next.Data)
		}
	})
}
