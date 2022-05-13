package stack

import "fmt"

var head *Node //the first node

// nest struct to define linkedList node
type Node struct {
	value int
	next  *Node
}

func StackUsingLinkedList() {
	head = &Node{}
}

// Remove value from the beginning of the list to simulate stack
func PopList() int {
	if head == nil {
		fmt.Print("Error remove from of the list")
	}
	valu := head.value
	fmt.Print("Popped element :", valu)

	head = head.next
	return valu
}

// Add value to the beginning of the list to simulate stack
func PushList(value int) {
	fmt.Print("Pushed element:", value)
	prevHead := head
	head = &Node{}
	head.value = value
	head.next = prevHead
}
