package test

import (
	"reflect"
	"testing"

	. "github.com/amtc131/algorithms/structure/linkedList/reverse"
)

func TestReverse(t *testing.T) {

	t.Run("Test 1: Reverse Linked List", func(t *testing.T) {
		var sll *MyLinkedList = GetLinkedList()
		//add new elements in linked list
		sll.Add(1)
		sll.Add(4)
		sll.Add(3)
		sll.Add(7)
		sll.Add(8)
		var want = sll.Head
		// Before Reverse
		// 1 4 3 7 8
		// Set new last node
		sll.Tail = sll.Reverse(sll.Head)
		// 8 7 3 4 1
		if !reflect.DeepEqual(sll.Tail, want) {
			t.Errorf("Reverse()= %#v; want= #%v", sll.Tail, want)
		}
	})
}
