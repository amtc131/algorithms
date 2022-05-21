package reverse

import "fmt"

type LinkNode struct {
	Data int
	Next *LinkNode
}

func GetLinkNode(data int) *LinkNode {
	return &LinkNode{Data: data, Next: nil}
}

type MyLinkedList struct {
	Head *LinkNode
	Tail *LinkNode
}

func GetLinkedList() *MyLinkedList {
	return &MyLinkedList{}
}

// Add new node at end of linked list
func (t *MyLinkedList) Add(data int) {
	//Created a Node
	var node *LinkNode = GetLinkNode(data)
	if t.Head == nil {
		t.Head = node
	} else {
		t.Tail.Next = node
	}
	//Set new last node
	t.Tail = node
}

func (t *MyLinkedList) Display() {
	var temp *LinkNode = t.Head
	if temp == nil {
		fmt.Println("Empty linked list")
	}
	// Iterating the linked list
	for temp != nil {
		//Display node value
		fmt.Print(" ", temp.Data)
		// Visit to next node
		temp = temp.Next
	}
}
