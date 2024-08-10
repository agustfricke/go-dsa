/*
Linked List is a linear data structure, in which elements are not stored at 
a contiguous location, rather they are linked using pointers. Linked List 
forms a series of connected nodes, where each node stores the data and 
the address of the next node.

              node                node               node
head --> [ data | next ] --> [ data | next ] --> [ data | next ] --> nil

*/
package linkedlist

import (
  "errors"
  "fmt"
)

// Node represents a node in the linked list
type Node struct {
  Data int
  Next *Node
}

// LinkedList represents the linked list itself
type LinkedList struct {
  Head *Node
}

// Insert adds a new node with the given data at the end of the list
func (ll *LinkedList) Insert(data int) {
  // create a new node with the data
  newNode := &Node{Data: data}
  // if the list is empty, set the head to the new node
  if ll.Head == nil {
    ll.Head = newNode
    // if the list is not empty, traverse the list until the end and add the new node
  } else {
    // current is a temporary pointer initialized to ll.Head, 
    // which points to the first node in the list.
    current := ll.Head
    // traverse the list and check if the current node has a next node
    for current.Next != nil {
      // if it doesn't, move to the next node
      current = current.Next
    }
    // Once the loop finishes, current will point to the last 
    // node in the list (whose Next is nil).
    // Set current.Next to newNode, effectively appending newNode 
    // to the end of the list.
    current.Next = newNode
  }
}

// Read traverses the list and returns the data at the given index
func (ll *LinkedList) Read(index int) (int, error) {
  // A pointer to traverse through the nodes of the linked list, 
  // initially set to the head.
  current := ll.Head
  // An integer to keep track of the current position in the list
  count := 0

  // For each node, check if the current position (count) 
  // matches the index provided
  for current != nil {
    if count == index {
      // If it matches, return the data of the node
      return current.Data, nil
    }
    // If it doesn’t match, move to the next node and increment count
    count++
    current = current.Next
  }
  // not found
  return 0, errors.New("index out of range")
}

// Delete removes the node at the given index
func (ll *LinkedList) Delete(index int) error {
  /*
    Check for Empty List:
    If ll.Head is nil, it means the list is empty.
    In this case, return an error indicating that the list is empty.
  */
  if ll.Head == nil {
    return errors.New("list is empty")
  }

  /*
    Deleting the Head Node:
    If index is 0, it means the node to delete is the head of the list.
    Set ll.Head to ll.Head.Next, effectively removing the current head node.
    Return nil to indicate that the operation was successful.
  */
  if index == 0 {
    ll.Head = ll.Head.Next
    return nil
  }

  /*
    Initialization for Traversal:
    current is initialized to ll.Head, pointing to the first node in the list.
    previous is initialized to an empty node, which will be used to keep track 
    of the node before current.
  */
  current := ll.Head
  previous := &Node{}

  /*
    Traversing the List:
    The for loop iterates from 0 up to index - 1.
    For each iteration, check if current.Next is nil. If it is, it means the index 
    is out of range (the list doesn’t have enough nodes).
    Update previous to be current, and move current to the next node 
    (current = current.Next).
  */
  for i := 0; i < index; i++ {
    if current.Next == nil {
      return errors.New("index out of range")
    }
    previous = current
    current = current.Next
  }

  /*
    Deleting the Node:
    Set previous.Next to current.Next, effectively skipping over current and 
    removing it from the list.
    Return nil to indicate that the operation was successful
  */
  previous.Next = current.Next
  return nil
  /*
    Example Linked List
    Let's use a linked list with the following nodes:
    Node1 (Value: 10) → Node2 (Value: 20) → Node3 (Value: 30) → Node4 (Value: 40)

    Case 1: Delete Node at Index 2
    Objective: Delete the node with value 30.

    Initial State:
      current points to Node1 (Value: 10).
      previous is initially an empty node or nil.

    Fist iteration 0 (i = 0):
      Check if current.Next is nil. It is not (Node2 exists).
      Update previous to current (Node1).
      Move current to current.Next (Node2).

    Second iteration 1 (i = 1):
      Check if current.Next is nil. It is not (Node3 exists).
      Update previous to current (Node2).
      Move current to current.Next (Node3).

    At the end of the loop:
      previous points to Node2 (Value: 20).
      current points to Node3 (Value: 30).

    Delete Operation:
      Set previous.Next to current.Next, effectively skipping Node3:
      previous.Next (Node2's Next) now points to current.Next (Node4).

    List after deletion:
      Node1 (10) → Node2 (20) → Node4 (40)


    Case 2: Delete Node at Index 0 (Head Node)
    Objective: Delete the head node (Node1 with value 10).

    Initial State:
      current points to Node1 (Value: 10).
      previous is initially an empty node or nil.

    Special Case Handling:
      If index == 0, the head node is removed:
      Set ll.Head to ll.Head.Next (Node2).

    List after deletion:
      Node2 (20) → Node3 (30) → Node4 (40)
  */
}

// PrintList prints all the elements in the list
func (ll *LinkedList) PrintList() {
  current := ll.Head
  for current != nil {
    fmt.Printf("%d -> ", current.Data)
    current = current.Next
  }
  fmt.Println("nil")
}
