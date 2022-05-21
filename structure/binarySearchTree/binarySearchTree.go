package binarySearchTree

import "fmt"

func Search(root, nodeToBeSearched *Node) bool {
	if root == nil {
		return false
	}
	if root.Data == nodeToBeSearched.Data {
		return true
	}

	var result bool = false

	if root.Data > nodeToBeSearched.Data {
		result = Search(root.Left, nodeToBeSearched)
	} else if root.Data < nodeToBeSearched.Data {
		result = Search(root.Right, nodeToBeSearched)
	}

	return result
}

func Insert(root *Node, nodeToBeInserted *Node) *Node {
	if root == nil {
		root = nodeToBeInserted
		return root
	}
	if root.Data > nodeToBeInserted.Data {
		if root.Left == nil {
			root.Left = nodeToBeInserted
		} else {
			Insert(root.Left, nodeToBeInserted)
		}
	} else if root.Data < nodeToBeInserted.Data {
		if root.Right == nil {
			root.Right = nodeToBeInserted
		} else {
			Insert(root.Right, nodeToBeInserted)
		}
	}
	return root
}

func InOrder(root *Node) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("%v ", root.Data)
		InOrder(root.Right)
	}
}
