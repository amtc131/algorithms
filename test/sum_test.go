package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestSum(t *testing.T) {
	tests := []struct {
		number []int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{3, 3}, 6},
		{[]int{3, 5, 3, 5, 3}, 19},
		{[]int{4, 2, 22, -10, 8}, 26},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.number), func(t *testing.T) {
			got := module.Sum(tc.number)
			if got != tc.want {
				t.Fatalf("Sum() = %v; want %v", got, tc.want)
			}
		})

	}

}
