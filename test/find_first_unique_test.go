package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestFindFirstUnique(t *testing.T) {
	tests := []struct {
		list []int
		want int
	}{
		{[]int{9, 2, 3, 2, 6, 6}, 9},
		{[]int{4, 4, 4, 3}, 3},
		{[]int{4, 5, 1, 2, 0, 4}, 5},
		{[]int{2, 54, 7, 2, 6, 54}, 7},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.list), func(t *testing.T) {
			got := module.FindFirstUnique(tc.list)
			if got != tc.want {
				t.Fatalf("FindFirstUnique() = %v; want %v", got, tc.want)
			}
		})

	}
}
