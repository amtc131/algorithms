package linkedlist

import "fmt"

// Doubly LinkedList
// In the doubly linkedList, Node has data and pointers to next node and previous node.
// You can iterate over linkedList either in forward or backward derections as it has pointers to prev node
// and next node. This is one of most used data structures.
//            Head                     Tail
//             |                        |
//               -->    -->   -->   -->   -->
//  null  <--  1 <--  7 <-- 6 <-- 5 <-- 2      null
//
//             Doubly Linked List

var tail *DoublyNode
var size int

type DoublyNode struct {
	Data int
	Next *DoublyNode
	Prev *DoublyNode
	Head *DoublyNode
}

func DoublyLinkedList() *DoublyNode {
	return &DoublyNode{}
}

func (n *DoublyNode) DoublyInsertFirts(data int) {
	newNode := &DoublyNode{}

	newNode.Data = data
	newNode.Next = n.Head
	newNode.Prev = nil
	if n.Head != nil {
		n.Head.Prev = newNode
	}
	n.Head = newNode
	if tail == nil {
		tail = newNode
	}
	size++
}

func (n *DoublyNode) DoublyInsertLast(data int) {
	newNode := &DoublyNode{}
	newNode.Data = data
	newNode.Next = nil
	newNode.Prev = tail
	if tail != nil {
		tail.Next = newNode
	}
	tail = newNode
	if n.Head == nil {
		n.Head = newNode
	}
	size++
}

// Used to delete node from start of doubly linked list
func (n *DoublyNode) DoublyDeleteFirst() *DoublyNode {
	if n.isEmpty() {
		fmt.Print("Doubly linked list is already empty")
		return nil
	}

	temp := n.Head
	n.Head = n.Head.Next
	n.Head.Prev = nil
	size--
	return temp
}

// Method to delete node from last of Doubly list
func (n *DoublyNode) DoublyDeleteLast() *DoublyNode {
	temp := tail

	tail = tail.Prev
	tail.Next = nil
	size--
	return temp
}

// Method to delete node after particular node
func (n *DoublyNode) DoublyDeleteAfter(after *DoublyNode) {
	temp := n.Head
	for ; temp.Next != nil && temp.Data != after.Data; temp = temp.Next {
	}
	if temp.Next != nil {
		temp.Next.Next.Prev = temp
	}
	temp.Next = temp.Next.Next
}

func (n *DoublyNode) isEmpty() bool {
	return (n.Head == nil)
}
