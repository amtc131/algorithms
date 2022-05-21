package reverse

// Reverse linked list node
func (t *MyLinkedList) Reverse(node *LinkNode) *LinkNode {
	if node == nil {
		return nil
	}

	//Recursively find next node
	var nextNode *LinkNode = t.Reverse(node.Next)
	if node == t.Tail {
		// Get last node
		//Set new head
		t.Head = node
	} else {
		// connect next node to current noce
		nextNode.Next = node
	}
	// unlink current node
	node.Next = nil
	return node
}
