package test

import (
	"reflect"
	"testing"

	linkedlist "github.com/amtc131/algorithms/structure/linkedList"
)

func TestDoublyLinkedList(t *testing.T) {
	list := linkedlist.DoublyLinkedList()

	t.Run("Test 1: DeleteFirst() list is empty", func(t *testing.T) {
		got := list.DoublyDeleteFirst()
		if got != nil {
			t.Errorf("DeleteFirst()= %v, want =%v", got, nil)
		}
	})

	list.DoublyInsertFirts(50)
	list.DoublyInsertFirts(60)
	list.DoublyInsertFirts(70)
	list.DoublyInsertFirts(10)

	t.Run("Test 2: DoublyInsertFirts() elements on the Node (Head --> tail)", func(t *testing.T) {
		want := []int{10, 70, 60, 50}
		got := []int{}
		currency := list.Head
		got = append(got, currency.Data)

		for currency.Next != nil {
			currency = currency.Next
			got = append(got, currency.Data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoublyInsertFirts() got: %v, want: %v", got, want)
		}
	})

	list.DoublyInsertLast(20)
	t.Run("Test 3: DoublyInsertLast(20) (Head --> tail)", func(t *testing.T) {
		want := []int{10, 70, 60, 50, 20}
		got := []int{}
		currency := list.Head
		got = append(got, currency.Data)

		for currency.Next != nil {
			currency = currency.Next
			got = append(got, currency.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoublyInsertLast() got: %v, want: %v", got, want)
		}
	})

	t.Run("Test 4: DeleteAfter(10)", func(t *testing.T) {
		want := []int{10, 60, 50, 20}
		got := []int{}

		newDoubliNode := linkedlist.DoublyLinkedList()

		newDoubliNode.Data = 10
		currency := list.Head
		got = append(got, currency.Data)

		list.DoublyDeleteAfter(newDoubliNode)

		for currency.Next != nil {
			currency = currency.Next
			got = append(got, currency.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("DeleteAfter() got: %v, want: %v", got, want)
		}
	})

	t.Run("Test 4: DeleteFirst()", func(t *testing.T) {
		want := []int{60, 50, 20}
		got := []int{}

		list.DoublyDeleteFirst()

		currency := list.Head
		got = append(got, currency.Data)

		for currency.Next != nil {
			currency = currency.Next
			got = append(got, currency.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("DeleteFirst() got: %v, want: %v", got, want)
		}
	})
	t.Run("Test 4: DeleteLast()", func(t *testing.T) {
		want := []int{60, 50}
		got := []int{}

		list.DoublyDeleteLast()

		currency := list.Head
		got = append(got, currency.Data)

		for currency.Next != nil {
			currency = currency.Next
			got = append(got, currency.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("DeleteLast() got: %v, want: %v", got, want)
		}
	})

}
