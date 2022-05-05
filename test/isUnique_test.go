package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"alex", true}, // a a l e x
		{"cat", true},
		{"alphabet", false},
		{"unique", false},
		{"o", true},
		{"oo", false},
		{"fizzbuzz", false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.word), func(t *testing.T) {
			got := module.IsUnique(tc.word)
			if got != tc.want {
				t.Fatalf("IsUnique() = %v; want %v", got, tc.want)
			}
		})

	}
}
