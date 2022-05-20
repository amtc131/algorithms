package test

import (
	"reflect"
	"testing"

	"github.com/amtc131/algorithms/module"
)

func Test(t *testing.T) {

	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{10, 20, 30, 40, 50}, []int{50, 40, 30, 20, 10}},
		{[]int{5, 10, 15, 20, 25}, []int{25, 20, 15, 10, 5}},
		{[]int{0, 2, 5, 9, 0}, []int{0, 9, 5, 2, 0}},
	}

	// 1 2 3 4 5 | 5 4 3 2 1
	// temp = arr[0]
	// arr[0]= arr[4]
	// arr[4] = temp
	for _, tc := range tests {
		t.Run("Test 1: ReverseArray", func(t *testing.T) {
			module.ReverseArray(tc.arr)
			got := tc.arr
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReverseArray(%v) = %v; want = %v", tc.arr, got, tc.want)
			}
		})
	}
}
