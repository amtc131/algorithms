package test

import (
	"fmt"
	"testing"

	"github.com/amtc131/algorithms/structure/binarySearchTree"
)

func TestBinarySearch(t *testing.T) {

	test := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, tc := range test {
		t.Run(fmt.Sprintf("Test BinarySearch(%v, %v); want = %v", tc.arr, tc.target, tc.want), func(t *testing.T) {
			got := binarySearchTree.BinarySearch(tc.arr, tc.target)

			if got != tc.want {
				t.Fatalf("BinarySearch(%v, %v) = %v; want = %v", tc.arr, tc.target, got, tc.want)
			}
		})
	}

}
