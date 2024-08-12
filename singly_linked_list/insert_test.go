package singlylinkedlist_test

import (
	"testing"

	singlylinkedlist "github.com/agustfricke/dsa-go/singly_linked_list"
)

func TestInsert(t *testing.T) {
	t.Run("Insert a node in empty list", func(t *testing.T) {
		ll := &singlylinkedlist.LinkedList{}
		ll.Insert(420)
		if ll.Head.Data != 420 {
			t.Fatalf("Expected head to be 420, got %d", ll.Head.Data)
		}
	})

	t.Run("Insert a node in non-empty list", func(t *testing.T) {
		ll := &singlylinkedlist.LinkedList{}
		ll.Insert(420)
		ll.Insert(69)
		if ll.Head.Data != 69 {
			t.Fatalf("Expected head to be 69, got %d", ll.Head.Data)
		}
		if ll.Head.Next.Data != 420 {
			t.Fatalf("Expected next node to be 420, got %d", ll.Head.Next.Data)
		}
	})
}
