package linkedlist

// LinkedList
// in singly linked list, Node has data and pointer to next node.
// It does not have pointer to the previous node. Last node's next points
// to null, so you can iterate, so you can iterate over linked
// list by using this conditions
//
//	Head
//   |
//   5 -> 6 -> 7 -> 1 -> 2 -> null
//

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
	Head *Node
}

func NewNodeLinkedList() *Node {
	return &Node{}
}

// For Inserting node at the start of linked list
func (n *Node) InsertFirst(data int) {
	newHead := &Node{}

	newHead.Data = data
	newHead.Next = n.Head
	n.Head = newHead
}

// Method for deleting node from start of linked list
func (n *Node) DeleteFirts() *Node {
	temp := n.Head
	n.Head = n.Head.Next
	return temp
}

// Method used to delete node after provided node
func (n *Node) DeleteAfter(after *Node) {
	temp := n.Head

	for temp.Next != nil && temp.Data != after.Data {
		temp = temp.Next
	}
	if temp.Next != nil {
		temp.Next = temp.Next.Next
	}
}

// Method used to insert at end of linkedList
func (n *Node) InsertLast(data int) {
	current := n.Head

	for ; current.Next != nil; current = current.Next {
	}

	newNode := &Node{}

	newNode.Data = data
	current.Next = newNode
}

// Method validate if the node is emty
func (n *Node) IsEmpty() bool {
	return n.Head == nil
}

func (n *Node) PrintLinkedList() {
	fmt.Println("Printing LinkedList (head --> last)")
	current := n.Head
	for current.Next != nil {
		current.displayNodeData()
		current = current.Next
	}
}

func (n *Node) displayNodeData() {
	fmt.Println("{ ", n.Data, " }")
}
