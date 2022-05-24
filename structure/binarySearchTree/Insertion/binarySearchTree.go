package insertion

import "fmt"

func (bs BinarySearchTree) Preorder(node *TreeNode) {
	if node != nil {
		// Display node value
		fmt.Print(" ", node.Data)
		//visit to left subtree
		bs.Preorder(node.Left)
		// visit to right subtree
		bs.Preorder(node.Right)
	}
}

func (bs BinarySearchTree) Inorder(node *TreeNode) {
	if node != nil {
		// visit to left subtree
		bs.Inorder(node.Left)
		//Display node value
		fmt.Print(" ", node.Data)
		// Visit to right subtree
		bs.Inorder(node.Right)
	}
}

func (bs BinarySearchTree) Postorder(node *TreeNode) {
	if node != nil {
		// visit to left subtree
		bs.Inorder(node.Left)
		// Visit to right subtree
		bs.Inorder(node.Right)
		//Display node value
		fmt.Print(" ", node.Data)
	}
}
