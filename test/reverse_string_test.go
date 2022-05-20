package test

import (
	"fmt"
	"testing"

	module "github.com/amtc131/algorithms/module"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"ABCDE", "EDCBA"},
		{"ca t", "t ac"},
		{"alphabet", "tebahpla"},
		{"654A321", "123A456"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("ReverseStr(%v)", tc.word), func(t *testing.T) {
			got := module.ReverseStr(tc.word)
			if got != tc.want {
				t.Fatalf("ReverseStr() = %v; want %v", got, tc.want)
			}
		})

	}

}
