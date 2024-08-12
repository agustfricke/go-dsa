package singlylinkedlist

import "errors"

func (ll *LinkedList) Delete() error {
  if ll.Head == nil {
    return errors.New("list is empty")
  }
  ll.Head = ll.Head.Next
  return nil
}
