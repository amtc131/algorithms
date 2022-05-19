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
		// push to stack
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		// pop in stack
		curren := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curren.Data)

		// we push Rigth first and Left next because we want to process
		// left first and right next
		if curren.Right != nil {
			// Push to stack Rigth
			stack = append(stack, curren.Right)
		}
		if curren.Left != nil {
			// Push to stack Left
			stack = append(stack, curren.Left)
		}
	}
}

// inorder traversal of binary tree
func (this BinaryTree) InOrder(node *Node) {
	if node != nil {
		// Visit left subtree
		this.InOrder(node.Left)
		// Print node value
		fmt.Print("  ", node.Data)
		// Visit right subtree
		this.InOrder(node.Right)
	}
}

// Check that given key exists in binary tree
func (this BinaryTree) findNode(node *Node,
	key int) bool {
	if node == nil {
		return false
	} else if node.Data == key {
		// Found the key node then return
		// 1 indicates node exists.
		return true
	} else if this.findNode(node.Left, key) ||
		this.findNode(node.Right, key) {
		return true
	}
	return false
}

func (this BinaryTree) BetweenInorder(node *Node, n1 int, n2 int, flag []bool) {
	if node != nil {
		// Visit left subtree
		this.BetweenInorder(node.Left, n1, n2, flag)
		if n1 == node.Data &&
			flag[0] == false {
			// We get first node n1
			flag[0] = true
		} else if n2 == node.Data && flag[1] == false {
			// We get second node n2
			flag[1] = true
		} else if flag[0] != flag[1] {
			this.status = true
			// Print intermediate inorder nodes
			fmt.Print("  ", node.Data)
		}
		// Visit right subtree
		this.BetweenInorder(node.Right, n1, n2, flag)
	}
}

func (this *BinaryTree) NodeInorder(n1, n2 int) {
	if this.Root == nil {
		return
	}
	if this.findNode(this.Root, n1) && this.findNode(this.Root, n2) {
		// When both node exists
		// flag are used to indicate node status.
		var flag = []bool{
			false,
			false,
		}
		this.status = false
		// Display node value
		fmt.Println("\n Inorder between (",
			n1, " ", n2, ") is")
		// Inorder between n1 and n2.
		this.BetweenInorder(this.Root, n1, n2, flag)
		if this.status == false {
			fmt.Println(" None ")
		}
	} else {
		fmt.Println(" \n Given node pair (",
			n1, ",", n2, ") are not exist ")
	}
}
