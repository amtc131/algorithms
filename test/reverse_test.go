package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"cat", "tac"},
		{"ca t", "t ac"},
		{"alphabet", "tebahpla"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.word), func(t *testing.T) {
			got := module.Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse() = %v; want %v", got, tc.want)
			}
		})

	}

}
