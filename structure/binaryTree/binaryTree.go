package binarytree

import "fmt"

// There are two standard solution to traversal of binary tree in inorder form.
// First is recursion and second is iterative approach using stack.
// In this post view both of solution.
// Before traversal of tree in in-order form that is important
// to know about what will be the result of inorder traversal of any tree.

func (b BinaryTree) Preorder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Data)
		b.Preorder(root.Left)
		b.Preorder(root.Right)
	}
}

// preorder traversal
// here the stack is initialized with the root node and inside the processing
// loop the node is visited first and then its right and left nodes are pushed
// to the stack in that order (if  they are empty).
func (b BinaryTree) PreorderInter(root *Node) {
	stack := []*Node{}

	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		curren := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curren.Data)
		if curren.Right != nil {
			stack = append(stack, curren.Right)
		}
		if curren.Left != nil {
			stack = append(stack, curren.Left)
		}
	}
}
