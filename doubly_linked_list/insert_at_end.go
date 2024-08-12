package doublylinkedlist
 
func (dll *DoublyLinkedList) InsertAtEnd(data int) {
  newNode := &Node{Data: data}
  if dll.Head == nil {
    dll.Head = newNode
    dll.Tail = newNode
  } else {
    newNode.Prev = dll.Tail
    dll.Tail.Next = newNode
    dll.Tail = newNode
  }
}
