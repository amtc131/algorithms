package binarySearchTree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func GetNode(data int) *Node {
	return &Node{Data: data, Left: nil, Right: nil}
}
