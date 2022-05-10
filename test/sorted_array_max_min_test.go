package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestMaxMin(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 1, 6, 2, 5, 3, 4}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.list), func(t *testing.T) {
			module.MaxMin(tc.list)
			for i, v := range tc.list {
				if v != tc.want[i] {
					t.Fatalf("TestMaxMin() = %v; want %v", v, tc.want)
				}
			}
		})

	}
}
