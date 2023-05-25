package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data string) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}

	if current == nil {
		fmt.Print("nil ")
	}

	fmt.Println()
}

func main() {
    reader := bufio.NewReader(os.Stdin)
	list := LinkedList{}
    for {
        var task string
        fmt.Print("Enter a new task: ")
        task, _ = reader.ReadString('\n')
        task = strings.TrimSpace(task)
        list.Insert(task)
            
        list.Display() 

        var exit string
        fmt.Print("Type y to quit: ")
        fmt.Scanf("%v\n", &exit)
        if(exit == "y") {
            break
        }

    }
}
