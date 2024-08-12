package doublylinkedlist

import (
	"errors"
)

func (dll *DoublyLinkedList) DeleteTail() error {
  if dll.Tail == nil {
    return errors.New("Empty list")
  } 

  if dll.Tail.Prev == nil {
    dll.Tail = nil
    dll.Head = nil
    return nil
  }

  dll.Tail = dll.Tail.Prev
  dll.Tail.Next = nil
  return nil
}
