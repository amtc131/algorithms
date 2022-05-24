package insertion

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func GetTreeNode(data int) *TreeNode {
	return &TreeNode{Data: data, Left: nil, Right: nil}
}

type BinarySearchTree struct {
	Root *TreeNode
}

func GetBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// 1 First inserted node is called root
// 2 Every node can be hold 2 child node. Which is called left and right child or called left and right subtree
// 3 Whitch contains at least one child is called parent node
// 4 When node not cantain any child nodes this is called leaft node
// 5 Node Key of left subtree can be less than or equals to every parent node key.
// 6 Node key of right subtree always be greater than or equals to every parent node key
// 5th and 6th point is very important to implement algortims. Try to understand its with this example
//            10                 10           10               10              10               10                   10
//          /   \                            /                /  \            /  \            /    \               /    \
//        4     15       ---->       ----> 4       ---->     4    15 ---->   4   15 ---->    4      15    ---->   4      15
//      /  \   /                                                              \               \    /            /  \    /
//     3    5 12                                                               5               5  12           3    5  12
//                           Add First     Add node         Add node         Add node         Add node        Add node
//                             Node	       element 4        elemtn 15        elemtn 5         elemtn 12        elemtn 3
//                                          4 < 10           15 > 10         5<10 5>4         12 > 10          3 < 10
//                                                                                            12 < 15          3 < 4
func (t *BinarySearchTree) AddNode(data int) {
	// create a new node
	var node *TreeNode = GetTreeNode(data)

	if t.Root == nil {
		//When adds a first node in bst
		t.Root = node
	} else {
		var find *TreeNode = t.Root
		// Add new node to proper position
		for find != nil {
			if find.Data >= data {
				if find.Left == nil {
					// When left child emty
					// So add new node here
					find.Left = node
					return
				} else {
					// otherwiese
					// visit left sub-tree
					find = find.Left
				}
			} else {
				if find.Right == nil {
					//When right child empty
					// So add new node here
					find.Right = node
					return
				} else {
					//visit right sub-tree
					find = find.Right
				}

			}
		}
	}
}
