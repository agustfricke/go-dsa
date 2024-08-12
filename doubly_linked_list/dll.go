package doublylinkedlist

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}
