package linkedlist_test

import (
  "testing"

  linkedlist "github.com/agustfricke/dsa-go/ds/linked_list"
)

func TestLinkedList_Insert(t *testing.T) {
  // new empty linked list
  ll := &linkedlist.LinkedList{}
  // insert some values
  ll.Insert(10)
  ll.Insert(20)
  ll.Insert(30)

  // the list should look like this: 10 -> 20 -> 30
  expected := []int{10, 20, 30}
  // initialize current with the first node
  current := ll.Head

  // traverse the list and check the values
  for i, v := range expected {
    // if the current node is nil, the list is not long enough
    if current == nil {
      t.Fatalf("expected node with value %d at index %d, but got nil", v, i)
    }
    // check the values, if they don't match, the test fails
    if current.Data != v {
      t.Errorf("expected %d, got %d", v, current.Data)
    }
    // move to the next node
    current = current.Next
  }

  // after traversing the list, the current node should be nil
  if current != nil {
    t.Error("expected the list to end, but it continues")
  }
}

func TestLinkedList_Read(t *testing.T) {
  // new empty linked list
  ll := &linkedlist.LinkedList{}
  // insert some values
  ll.Insert(10)
  ll.Insert(20)
  ll.Insert(30)

  // slice of test cases
  tests := []struct {
    index    int
    expected int
    err      bool
  }{
    {0, 10, false},
    {1, 20, false},
    {2, 30, false},
    {3, 0, true},
  }

  // iterate over the test cases
  for _, test := range tests {
    // read the value at the given index
    result, err := ll.Read(test.index)
    // check if the error is as expected
    if (err != nil) != test.err {
      t.Errorf("expected error: %v, got: %v", test.err, err)
    }
    // check if the result is as expected
    if result != test.expected {
      t.Errorf("expected %d, got %d", test.expected, result)
    }
  }
}

func TestLinkedList_Delete(t *testing.T) {
  // new empty linked list
  ll := &linkedlist.LinkedList{}
  // insert some values
  ll.Insert(10)
  ll.Insert(20)
  ll.Insert(30)

  err := ll.Delete(1) // Deleting the node with value 20
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }

  // Verify the linked list after deletion
  expected := []int{10, 30}
  current := ll.Head

  for i, v := range expected {
    if current == nil {
      t.Fatalf("expected node with value %d at index %d, but got nil", v, i)
    }
    if current.Data != v {
      t.Errorf("expected %d, got %d", v, current.Data)
    }
    current = current.Next
  }

  if current != nil {
    t.Error("expected the list to end, but it continues")
  }

  // Test deleting head
  err = ll.Delete(0)
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }
  if ll.Head.Data != 30 {
    t.Errorf("expected head to be 30, got %d", ll.Head.Data)
  }

  // Test deleting last remaining element
  err = ll.Delete(0)
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }
  if ll.Head != nil {
    t.Error("expected the list to be empty, but it's not")
  }
}

func TestLinkedList_DeleteFromEmptyList(t *testing.T) {
  ll := &linkedlist.LinkedList{}
  err := ll.Delete(0)
  if err == nil {
    t.Error("expected an error when deleting from an empty list, but got none")
  }
}

func TestLinkedList_ReadFromEmptyList(t *testing.T) {
  ll := &linkedlist.LinkedList{}
  _, err := ll.Read(0)
  if err == nil {
    t.Error("expected an error when reading from an empty list, but got none")
  }
}
