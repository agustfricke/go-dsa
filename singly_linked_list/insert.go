package singlylinkedlist

func (ll *LinkedList) Insert(data int) {
  newNode := &Node{Data: data}
  if ll.Head == nil {
    ll.Head = newNode
  } else {
    newNode.Next = ll.Head
    ll.Head = newNode
  }
}
