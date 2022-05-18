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

}
