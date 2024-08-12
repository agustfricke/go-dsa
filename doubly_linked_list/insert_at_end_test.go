package doublylinkedlist_test

import (
	"testing"

	doublylinkedlist "github.com/agustfricke/dsa-go/doubly_linked_list"
)

func TestInsertAtEnd(t *testing.T) {
  t.Run("insert one node", func(t *testing.T) {
    dll := doublylinkedlist.DoublyLinkedList{}
    dll.InsertAtEnd(73)
    if dll.Head.Data != 73 {
      t.Fatalf("Expected head to be 73, got %d", dll.Head.Data)
    }
    if dll.Tail.Data != 73 {
      t.Fatalf("Expected tail to be 73, got %d", dll.Tail.Data)
    }
    if dll.Tail.Prev != nil {
      t.Fatalf("Expected tail.prev to be nil, got %v", dll.Tail.Prev)
    }
    if dll.Head.Next != nil {
      t.Fatalf("Expected head.next to be nil, got %v", dll.Head.Next)
    }
    if dll.Head.Prev != nil {
      t.Fatalf("Expected head.prev to be nil, got %v", dll.Head.Prev)
    }
    if dll.Tail.Next != nil {
      t.Fatalf("Expected tail.next to be nil, got %v", dll.Tail.Next)
    }
  })

  t.Run("insert multiples nodes", func(t *testing.T) {
    dll := doublylinkedlist.DoublyLinkedList{}
    dll.InsertAtEnd(73)
    dll.InsertAtEnd(420)
    dll.InsertAtEnd(69)
    if dll.Head.Data != 73 {
      t.Fatalf("Expected head to be 73, got %d", dll.Head.Data)
    }
    if dll.Tail.Data != 69 {
      t.Fatalf("Expected tail to be 69, got %d", dll.Tail.Data)
    }
    if dll.Tail.Prev.Data != 420 {
      t.Fatalf("Expected tail.prev to be 420, got %v", dll.Tail.Next)
    }
    if dll.Head.Next.Data != 420 {
      t.Fatalf("Expected head.next to be 420, got %v", dll.Head.Next)
    }
    if dll.Head.Prev != nil {
      t.Fatalf("Expected head.prev to be nil, got %v", dll.Head.Prev)
    }
    if dll.Tail.Next != nil {
      t.Fatalf("Expected tail.next to be nil, got %v", dll.Tail.Next)
    }
  })
}
