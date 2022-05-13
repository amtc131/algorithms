package test

import (
	"testing"

	stack "github.com/amtc131/algorithms/structure/stack"
)

func TestMyStack(t *testing.T) {
	stack.MyStack(5)

	t.Run("Test 1: Stack is empty", func(t *testing.T) {
		got := stack.Pop()
		if got != -1 {
			t.Errorf("Pop() = %v, Got: %v", -1, got)
		}
	})
	
	t.Run("Test 2: Push "func (t *testing.T){})




















	})
}
