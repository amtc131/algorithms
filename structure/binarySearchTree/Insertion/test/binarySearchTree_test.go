package test

import (
	"os"
	"testing"

	. "github.com/amtc131/algorithms/structure/binarySearchTree/Insertion"
	"github.com/amtc131/algorithms/structure/utils"
)

func TestBinarySearchTree(t *testing.T) {
	var tree *BinarySearchTree = GetBinarySearchTree()

	tree.AddNode(10)
	tree.AddNode(4)
	tree.AddNode(3)
	tree.AddNode(5)
	tree.AddNode(15)
	tree.AddNode(12)

	t.Run("Test 1: Preoder(tree)", func(t *testing.T) {
		want := " 10 4 3 5 15 12"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.Preorder(tree.Root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Preorder(tree) = %q; want %q", got, want)
		}
	})

	t.Run("Test 1: Inorder(tree)", func(t *testing.T) {
		want := " 3 4 5 10 12 15"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.Inorder(tree.Root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Inorder(tree) = %q; want %q", got, want)
		}
	})
	t.Run("Test 1: Postorder(tree)", func(t *testing.T) {
		want := " 3 4 5 12 15 10"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		tree.Postorder(tree.Root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Postorder(tree) = %q; want %q", got, want)
		}
	})
}
