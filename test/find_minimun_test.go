package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestFindMinimum(t *testing.T) {
	tests := []struct {
		list []int
		want int
	}{
		{[]int{9, 3, 2, 6}, 2},
		{[]int{10, 1, 0, 5}, 0},
		{[]int{6, 1, 3, 8}, 1},
		{[]int{6, 7, 8, 20, 30, 40, 100, 9, 16, 19, 25, 35, 45, 21, 13, 11, 18, 15, 10, 50, 5}, 5},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.list), func(t *testing.T) {
			got := module.FindMinimum(tc.list)
			if got != tc.want {
				t.Fatalf("FindMinimum() = %v; want %v", got, tc.want)
			}
		})

	}
}
