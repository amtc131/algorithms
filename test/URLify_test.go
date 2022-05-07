package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestUrlify(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"Mr John Smith", "Mr%20John%20Smith"},
		{"Alex The Test", "Alex%20The%20Test"},
		{"Alex The ", "Alex%20The%20"},
		{"Not", "Not"},
		{"this is a test of space", "this%20is%20a%20test%20of%20space"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.word), func(t *testing.T) {
			got := module.Urlify(tc.word)
			if got != tc.want {
				t.Fatalf("Urlify() = [%v]; want [%v]", got, tc.want)
			}
		})

	}

}
