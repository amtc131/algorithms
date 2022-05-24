package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	. "github.com/amtc131/algorithms/structure/binarySearchTree"
	"github.com/amtc131/algorithms/structure/utils"
)

func TestBinarySearchTree(t *testing.T) {

	//            40
	//          /     \
	//        20      60
	//       /  \    /  \
	//      10  30  50  70
	//     /		      \
	//    5           55

	root := GetNode(40)
	root.Left = GetNode(20)
	root.Right = GetNode(60)

	root.Left.Left = GetNode(10)
	root.Left.Right = GetNode(30)
	root.Right.Left = GetNode(50)
	root.Right.Right = GetNode(70)

	root.Left.Left.Left = GetNode(5)
	root.Right.Left.Right = GetNode(55)

	t.Run("Test 1: Insert() Nodes", func(t *testing.T) {

		rNode := GetNode(40)
		node20 := GetNode(20)
		node10 := GetNode(10)
		node30 := GetNode(30)
		node60 := GetNode(60)
		node50 := GetNode(50)
		node70 := GetNode(70)
		node5 := GetNode(5)
		node55 := GetNode(55)

		Insert(nil, rNode)
		Insert(rNode, node20)
		Insert(rNode, node10)
		Insert(rNode, node30)
		Insert(rNode, node60)
		Insert(rNode, node50)
		Insert(rNode, node70)
		Insert(rNode, node5)
		Insert(rNode, node55)

		if !reflect.DeepEqual(rNode, root) {
			t.Errorf("Insert(); %+v; want= %+v", rNode, root)
		}
	})

	t.Run(fmt.Sprintf("Test 2: Search(%v) searching for node 55", 55), func(t *testing.T) {
		var node *Node = GetNode(55)
		got := Search(root, node)
		if got != true {
			t.Errorf("Search(55)= %v; want=%v", got, true)
		}
	})

	//            40
	//          /     \
	//        20      60
	//       /  \    /  \
	//      10  30  50  70
	//     /  \		    \
	//    5   13      55


	root.Left.Left.Right = GetNode(13)

	t.Run(fmt.Sprintf("Test 3: Insert(%v) ", 13), func(t *testing.T) {
		rNode := GetNode(40)
		node20 := GetNode(20)
		node10 := GetNode(10)
		node30 := GetNode(30)
		node60 := GetNode(60)
		node50 := GetNode(50)
		node70 := GetNode(70)
		node5 := GetNode(5)
		node55 := GetNode(55)
		node13 := GetNode(13)
		Insert(nil, rNode)
		Insert(rNode, node20)
		Insert(rNode, node10)
		Insert(rNode, node30)
		Insert(rNode, node60)
		Insert(rNode, node50)
		Insert(rNode, node70)
		Insert(rNode, node5)
		Insert(rNode, node55)

		got := Insert(rNode, node13)
		if !reflect.DeepEqual(got, root) {
			t.Errorf("Insert(root, 13)= %v; want=%v", got, root)
		}
	})

	t.Run(fmt.Sprintf("Test 2: InOrder() traversal of binary tree after adding %v", 13), func(t *testing.T) {
		want := "5 10 13 20 30 40 50 55 60 70 "
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()
		InOrder(root)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("InOrder() = %q; want %q", got, want)
		}
	})

}
