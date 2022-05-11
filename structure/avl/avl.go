package avl

type Node struct {
	Key         int
	Height      int
	Left, Right *Node
}

// NewTree create a new AVL tree
func NewTree() *Node {
	return nil
}

func Get(root *Node, key int) *Node {
	if root != nil {
		return nil
	}

	if root.Key == key {
		return root
	} else if root.Key < key {
		root = root.Right
	} else {
		root = root.Left
	}
	return Get(root, key)
}

func Insert(root **Node, key int) {
	if *root == nil {
		*root = &Node{
			Key:    key,
			Height: 1,
		}
		return
	}
	if (*root).Key < key {
		Insert(&(*root).Right, key)
	} else {
		Insert(&(*root).Left, key)
	}

	// Update height
	(*root).Height = height(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 {
		bFactor = balanceFactor((*root).Left)
		if bFactor == 1 { //LL
			llRotation(root)
		} else if bFactor == -1 { // LR
			lrRotation(root)
		}
	} else if bFactor == -2 { //R
		bFactor = balanceFactor((*root).Right)
		if bFactor == 1 { // RL
			rlRotation(root)
		} else if bFactor == -1 { // RR
			rrRotation(root)
		}
	}
}

func Delete(root **Node, key int) {
	if root == nil {
		return
	}
	if (*root).Key < key {
		Delete(&(*root).Right, key)
	} else if (*root).Key > key {
		Delete(&(*root).Left, key)
	} else {
		// 3. cases
		// 1. No Child
		// 2. With One Child
		// 3. With Two Child
		if (*root).Left == nil && (*root).Right == nil {
			*root = nil
		} else if (*root).Left == nil {
			*root = (*root).Right
		} else if (*root).Right == nil {
			*root = (*root).Left
		} else {
			minVal := min((*root).Right)
			(*root).Key = minVal
			Delete(root, minVal)
		}
		return
	}

	//update height
	(*root).Height = height(*root)
	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		switch balanceFactor((*root).Left) {
		case 1: // ll
			llRotation(root)
		case 2: // LR
			lrRotation(root)
		case 3: // ll OR LR
			llRotation(root)
		}
	} else if bFactor == -2 {
		switch balanceFactor((*root).Right) {
		case 1: // ll
			rlRotation(root)
		case 2: // LR
			rrRotation(root)
		case 3: // ll OR LR
			rrRotation(root)
		}
	}
}

// rotations
// 1. LL
// 2. LR
// 3. RR
// 4. RL
func llRotation(root **Node) {
	b := (*root).Left

	br := b.Right
	b.Right = *root
	(*root).Left = br
	(*root).Height = height(*root)
	b.Height = height(b)
	*root = b
}

func lrRotation(root **Node) {
	c := (*root).Left.Right
	cl := c.Left
	cr := c.Right

	c.Left = (*root).Left
	c.Right = (*root)
	c.Left.Right = cl

	(*root).Left = cr

	(*root).Height = height(*root)
	c.Left.Height = height(c.Left)
	c.Height = height(c)

	*root = c
}

func rrRotation(root **Node) {
	b := (*root).Right
	bl := b.Left
	b.Left = *root

	(*root).Right = bl
	(*root).Height = height(*root)
	b.Height = height(b)
	*root = b
}

func rlRotation(root **Node) {
	c := (*root).Right.Left
	cl := c.Left
	cr := c.Right

	c.Right = (*root).Right
	c.Right.Left = cr
	c.Left = *root
	(*root).Right = cl

	(*root).Height = height(*root)
	c.Right.Height = height(c.Right)
	c.Height = height(c)
	*root = c
}

// balancefoctor : -ve balance factor means subtree root is heavy toward left
// and +ve balance factor means subtree root is heavy toward right side
func balanceFactor(root *Node) int {
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	var leftHeight, rightHeight int

	if root.Left != nil {
		leftHeight = root.Left.Height
	}

	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}
	return 1 + max
}

func min(root *Node) int {
	if root.Left == nil {
		return root.Key
	}
	return min(root.Left)
}
