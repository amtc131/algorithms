package test

import (
	"bytes"
	"io"
	"os"
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

	tests := []struct {
		element int
		want    string
	}{
		{100, "Pushed element:100"},
		{90, "Pushed element:90"},
		{10, "Pushed element:10"},
		{50, "Pushed element:50"},
	}

	for _, tc := range tests {
		t.Run("Test 2: Push ", func(t *testing.T) {
			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer

			defer func() {
				os.Stdout = osStdout
			}()

			stack.Push(tc.element)

			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, testStdout)
			got := buf.String()
			if got != tc.want {
				t.Fatalf("Push(%d) = %q; want %q", tc.element, got, tc.want)
			}
		})
	}
	testsPop := []struct {
		want string
	}{
		{"Popped element :50"},
		{"Popped element :10"},
		{"Popped element :90"},
		{"Popped element :100"},
		{"Stack is empty!"},
	}
	for _, tc := range testsPop {
		t.Run("Test 3: Stack Pop", func(t *testing.T) {
			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer

			defer func() {
				os.Stdout = osStdout
			}()

			stack.Pop()

			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, testStdout)
			got := buf.String()
			if got != tc.want {
				t.Fatalf("Pop() = %q; want %q", got, tc.want)
			}
		})
	}

}
