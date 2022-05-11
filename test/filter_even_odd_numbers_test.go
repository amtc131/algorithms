package test

import (
	"fmt"
	"testing"

	"github.com/amtc131/algorithms/module"
)

func TestFilterEvenOddNumbers(t *testing.T) {

	tests := []struct {
		list   []int
		flt    func(int) bool
		events []int
		odds   []int
	}{
		{[]int{1, 2, 3, 4, 5, 7}, module.IsEvent, []int{2, 4}, []int{1, 3, 5, 7}},
		{[]int{2, 4, 6, 8, 10, 12}, module.IsEvent, []int{2, 4, 6, 8, 10, 12}, []int{}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.list), func(t *testing.T) {
			event, odd := module.Filter(tc.list, tc.flt)
			for i, e := range event {
				if e != tc.events[i] {
					t.Fatalf("Filter() = %v; want %v", e, tc.events)
				}
			}
			for i, e := range odd {
				if e != tc.odds[i] {
					t.Fatalf("Filter() = %v; want %v", e, tc.events)
				}
			}
		})
	}
}
