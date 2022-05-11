package test

import (
	"testing"

	"github.com/amtc131/algorithms/structure/avl"
)

func TestInsert(t *testing.T) {
	t.Run("LLRotation", func(t *testing.T) {
		root := avl.NewTree()
		avl.Insert(&root, 5)
		avl.Insert(&root, 4)
		avl.Insert(&root, 3)

		if root.Key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}
		if root.Left.Key != 3 {
			t.Errorf("left child should be = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}
		if root.Right.Key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}
	})
}
