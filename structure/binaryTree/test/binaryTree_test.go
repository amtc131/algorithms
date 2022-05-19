package test

import (
	"os"
	"testing"

	. "github.com/amtc131/algorithms/structure/binaryTree"
	"github.com/amtc131/algorithms/structure/utils"
)

func TestBinaryTree(t *testing.T) {
	rootNode := NewNode(30)
	nodeTwenty := NewNode(50)
	nodeTen := NewNode(60)
	nodeThirtyn := NewNode(80)
	nodeSixtyn := NewNode(90)
	nodeFivetyn := NewNode(100)
	nodeSeventy := NewNode(110)

	rootNode.Left = nodeTwenty
	rootNode.Right = nodeSixtyn

	nodeTwenty.Left = nodeTen
	nodeTwenty.Right = nodeThirtyn

	nodeSixtyn.Left = nodeFivetyn
	nodeSixtyn.Right = nodeSeventy
	t.Run("Test 1: Create Binary Tree <Preorder>", func(t *testing.T) {
		want := "30 50 60 80 90 100 110 "
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()
		b := GetBinaryTree()
		b.Preorder(rootNode)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Preorder(1) = %q; want %q", got, want)
		}
	})

	t.Run("Test 2: Create Binary Tree <PreorderInter>", func(t *testing.T) {
		want := "30 50 60 80 90 100 110 "
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()
		b := GetBinaryTree()
		b.PreorderInter(rootNode)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("PreorderInter(1) = %q; want %q", got, want)
		}
	})

	t.Run("Test 3: Display preorder view of binary tree", func(t *testing.T) {
		want := "17 11 1 39 40 21 10 2 "
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()
		/*
					Construct Binary Tree
					-------------
					      17
			        	/    \
					  11      21
					 / \     /  \
					1   39  10   2
					   /
					 40

		*/

		tree := GetBinaryTree()

		// Add Tree Node
		tree.Root = NewNode(17)
		tree.Root.Left = NewNode(11)
		tree.Root.Right = NewNode(21)
		tree.Root.Right.Right = NewNode(2)
		tree.Root.Right.Left = NewNode(10)
		tree.Root.Left.Left = NewNode(1)
		tree.Root.Left.Right = NewNode(39)
		tree.Root.Left.Right.Left = NewNode(40)

		tree.Preorder(tree.Root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Preorder() = %q; want %q", got, want)
		}
	})

	t.Run("Test 4: Display preorder view of binary tree", func(t *testing.T) {
		want := "15 24 35 54 62 13 "
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()
		/*
					Construct Binary Tree
					-------------
					      15
			        	/    \
					  24      54
					 /       /  \
					35     62   13
		*/

		tree := GetBinaryTree()

		// Add Tree Node
		tree.Root = NewNode(15)
		tree.Root.Left = NewNode(24)
		tree.Root.Right = NewNode(54)
		tree.Root.Right.Right = NewNode(13)
		tree.Root.Right.Left = NewNode(62)
		tree.Root.Left.Left = NewNode(35)

		tree.Preorder(tree.Root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Preorder() = %q; want %q", got, want)
		}
	})

	var tree *BinaryTree = GetBinaryTree()
	/*
				 4
				/ \
			   /   \
			  12    7
			 / \     \
			2   3     1
			   / \   /
			  6   8  5
			 /     \
			9       10
		 -----------------
		  Constructing binary tree
	*/
	tree.Root = NewNode(4)
	tree.Root.Left = NewNode(12)
	tree.Root.Left.Right = NewNode(3)
	tree.Root.Left.Right.Left = NewNode(6)
	tree.Root.Left.Right.Left.Left = NewNode(9)
	tree.Root.Left.Right.Right = NewNode(8)
	tree.Root.Left.Right.Right.Right = NewNode(10)
	tree.Root.Left.Left = NewNode(2)
	tree.Root.Right = NewNode(7)
	tree.Root.Right.Right = NewNode(1)
	tree.Root.Right.Right.Left = NewNode(5)

	t.Run("Test 5: case A InOrder", func(t *testing.T) {
		want := "  2  12  9  6  3  8  10  4  7  5  1"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.InOrder(tree.Root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})

	t.Run("Test 6: case B nodeInorder(2,10)", func(t *testing.T) {
		want := "\n Inorder between ( 2   10 ) is\n  12  9  6  3  8 None \n"
		//         ↑                  ↑
		//          ------------------
		//           12 9 6 3 8
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.NodeInorder(2, 10)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})

	t.Run("Test 7: case C nodeInorder(8,5)", func(t *testing.T) {
		want := "\n Inorder between ( 8   5 ) is\n  10  4  7 None \n"
		//
		//          10 4 7
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.NodeInorder(8, 5)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})

	t.Run("Test 8: case D nodeInorder(1,6)", func(t *testing.T) {
		want := "\n Inorder between ( 1   6 ) is\n  3  8  10  4  7  5 None \n"
		//
		//          3  8  10  4  7  5
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.NodeInorder(1, 6)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})

	t.Run("Test 8: case E nodeInorder(3,6)", func(t *testing.T) {
		want := "\n Inorder between ( 3   6 ) is\n None \n"
		//
		//          3  8  10  4  7  5
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.NodeInorder(3, 6)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})

	t.Run("Test 9: case F nodeInorder(3,11)", func(t *testing.T) {
		want := " \n Given node pair ( 3 , 11 ) are not exist \n"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.NodeInorder(3, 11)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})
}
