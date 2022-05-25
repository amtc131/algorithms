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

	t.Run("Test 2: Inorder(tree)", func(t *testing.T) {
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
	t.Run("Test 3: Postorder(tree)", func(t *testing.T) {
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

  t.Run("Test 4 Alternate Sum", func(t *testing.T){
    //            5    
    //          /   \
    //         /     \
    //        3       19
    //      /   \    /  \
    //     2     4  8    31
    //             / \  / \ 
    //            7  15 25 50
    //
    // add tree node
    var b *BinarySearchTree = GetBinarySearchTree()

      b.AddNode(5)
      b.AddNode(3)
      b.AddNode(19)
      b.AddNode(2)
      b.AddNode(4)
      b.AddNode(8)
      b.AddNode(31)
      b.AddNode(7)
      b.AddNode(25)
      b.AddNode(15)
      b.AddNode(50)

      want := 34

      got := b.AlternateLeafSum()
      if got != want {
        t.Fatalf("AlternateLeafSum() = %d; want= %d", got, want)
      }

  })

}
