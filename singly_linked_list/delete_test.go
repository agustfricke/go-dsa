package singlylinkedlist_test

import (
	"testing"

	singlylinkedlist "github.com/agustfricke/dsa-go/singly_linked_list"
)

func TestDelete(t *testing.T) {
	t.Run("Delete empty list", func(t *testing.T) {
		ll := &singlylinkedlist.LinkedList{}
		err := ll.Delete()
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
	t.Run("Delete non-empty list", func(t *testing.T) {
		ll := &singlylinkedlist.LinkedList{}
		ll.Insert(420)
		ll.Insert(69)
		ll.Insert(73)
		err := ll.Delete()
		if err != nil {
			t.Fatalf("expected error to be nil, got %v", err)
		}
		if ll.Head.Data != 69 {
			t.Fatalf("expected head to be 69, got %d", ll.Head.Data)
		}
		if ll.Head.Next.Data != 420 {
			t.Fatalf("expected next node to be 420, got %d", ll.Head.Next.Data)
		}
	})
}
