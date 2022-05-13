package test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/amtc131/algorithms/structure/stack"
)

func TestStackUsingLinkedList(t *testing.T) {
	stack.StackUsingLinkedList()

	tests := []struct {
		element int
		want    string
	}{
		{200, "Pushed element:200"},
		{150, "Pushed element:150"},
		{80, "Pushed element:80"},
		{90, "Pushed element:90"},
		{600, "Pushed element:600"},
		{175, "Pushed element:175"},
	}

	for _, tc := range tests {
		t.Run("Test 1: Push elements to stack", func(t *testing.T) {

			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer

			defer func() {
				os.Stdout = osStdout
			}()

			stack.PushList(tc.element)

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
		{"Popped element :175"},
		{"Popped element :600"},
	}
	for _, tc := range testsPop {
		t.Run("Test 2: Removed elements from LinkedList", func(t *testing.T) {

			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer

			defer func() {
				os.Stdout = osStdout
			}()

			stack.PopList()
			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, testStdout)
			got := buf.String()
			if got != tc.want {
				t.Fatalf("Pop() = %q; want %q", got, tc.want)
			}
		})
	}

	testsPush := []struct {
		element int
		want    string
	}{
		{100, "Pushed element:100"},
	}

	for _, tc := range testsPush {
		t.Run("Test 3: Push element 100 to LinKedList", func(t *testing.T) {
			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer

			defer func() {
				os.Stdout = osStdout
			}()

			stack.PushList(tc.element)
			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, testStdout)
			got := buf.String()
			if got != tc.want {
				t.Fatalf("Push() = %q; want %q", got, tc.want)
			}
		})
	}

	testsPop2 := []struct {
		element int
		want    string
	}{
		{100, "Popped element :100"},
	}

	for _, tc := range testsPop2 {
		t.Run("Test 3: Removed element 100 from LinKedList", func(t *testing.T) {
			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer

			defer func() {
				os.Stdout = osStdout
			}()

			stack.PopList()
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
