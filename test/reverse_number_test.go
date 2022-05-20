package test

import (
	"fmt"
	"testing"

	"github.com/amtc131/algorithms/module"
)

func TestReverseNumber(t *testing.T) {

	tests := []struct {
		number int
		want   int
	}{
		{12345, 54321},
		{78942, 24987},
		{1020, 201},
		{-28, -82},
		{10, 1},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("Test 1: ReverseNumber(%v)", tc.number), func(t *testing.T) {
			got := module.ReverseNumber(tc.number)
			if got != tc.want {
				t.Errorf("ReverseNumber(%v) = %v; want= %v", tc.number, got, tc.want)
			}
		})
	}
}
