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

func (bs *BinarySearchTree) LeafSum(node *TreeNode) int {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			// Case A
			// When node is leaf node.
			// Change status
			bs.alternate = !bs.alternate
			// Check node is alternate or not
			if bs.alternate {
				// When get alternate node
				return node.Data
			}
		} else {
			// Case B
			// When node is internal
			// Visit left and right subtree and
			// find alternat node.
			return bs.LeafSum(node.Left) + bs.LeafSum(node.Right)
		}
	}
	return 0
}

func (bs *BinarySearchTree) AlternateLeafSum() int {
	bs.alternate = false
	return bs.LeafSum(bs.Root)
}
