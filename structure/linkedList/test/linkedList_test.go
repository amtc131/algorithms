package test

import (
	"reflect"
	"testing"

	list "github.com/amtc131/algorithms/structure/linkedList"
)

func TestLinkedList(t *testing.T) {
	node := list.NewNodeLinkedList()

	t.Run("Test 1: Linked List is empty", func(t *testing.T) {
		got := node.IsEmpty()
		if got != true {
			t.Errorf("got: %v, want: %v", got, true)
		}
	})

	node.InsertFirst(50)
	node.InsertFirst(60)
	node.InsertFirst(70)
	node.InsertFirst(10)

	t.Run("Test 2: InsertFirst()", func(t *testing.T) {
		want := []int{10, 70, 60, 50}
		got := []int{}
		current := node.Head
		got = append(got, current.Data)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	node.InsertLast(20)

	t.Run("Test 3: InsertLast()", func(t *testing.T) {
		want := []int{10, 70, 60, 50, 20}
		got := []int{}

		current := node.Head
		got = append(got, current.Data)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test 4: Delete node after Node 60", func(t *testing.T) {
		want := []int{10, 70, 60, 20}
		got := []int{}

		deleteNode := list.NewNodeLinkedList()
		deleteNode.Data = 60
		node.DeleteAfter(deleteNode)

		current := node.Head

		got = append(got, current.Data)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test 5: Delete node first", func(t *testing.T) {
		want := []int{70, 60, 20}
		got := []int{}

		node.DeleteFirts()

		current := node.Head

		got = append(got, current.Data)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
