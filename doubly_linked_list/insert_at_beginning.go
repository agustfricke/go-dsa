package doublylinkedlist

func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{Data: data}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
}
