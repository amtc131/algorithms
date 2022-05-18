package binarytree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

type BinaryTree struct {
	Root *Node
}

func GetBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}
