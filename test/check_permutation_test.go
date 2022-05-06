package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestCheckPermutation(t *testing.T) {
	tests := []struct {
		worldOne string
		worldTwo string
		want     bool
	}{
		{"test", "ttew", false},
		{"geeksforgeeks", "forgeeksgeeks", true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v, %v)", tc.worldOne, tc.worldTwo), func(t *testing.T) {
			got := module.CheckPermutation(tc.worldOne, tc.worldTwo)
			if got != tc.want {
				t.Fatalf("CheckPermutation() = %v; want= %v", got, tc.want)
			}
		})
	}
}
