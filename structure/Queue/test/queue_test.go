package test

import (
	"fmt"
	"os"
	"testing"

	queue "github.com/amtc131/algorithms/structure/Queue"
	"github.com/amtc131/algorithms/structure/utils"
)

func TestQueue(t *testing.T) {
	queue.NewQueue(6)

	t.Run("Test 1: Enqueue(1) Add to the queue", func(t *testing.T) {
		want := "1 added to the queue"

		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		queue.Enqueue(1)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Enqueue(1) = %q; want %q", got, want)
		}
	})

	t.Run("Test 2: Dequeue() 1 removed from the queue", func(t *testing.T) {
		want := "1 removed from the queue"

		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		queue.Dequeue()

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Dequeue() = %q; want %q", got, want)
		}
	})

	tests := []struct {
		num  int
		want string
	}{
		{30, "30 added to the queue"},
		{44, "44 added to the queue"},
		{32, "32 added to the queue"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Test 3 Enqueue(%v) Add to the queue", tc.num), func(t *testing.T) {
			testStdout, writer, osStdout := utils.NewStDoutRestore(t)

			defer func() {
				os.Stdout = osStdout
			}()

			queue.Enqueue(tc.num)

			got := utils.Stdout(writer, testStdout)

			if got != tc.want {
				t.Fatalf("Enqueue(%q) = %q; want %q", tc.num, got, tc.want)
			}
		})
	}

	t.Run("Test 4: Dequeue Removed element 30 dequeue", func(t *testing.T) {
		want := "30 removed from the queue"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		queue.Dequeue()

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Dequeue() = %q; want %q", got, want)
		}
	})

	t.Run("Test 5: Enqueue() added to element 98 dequeue", func(t *testing.T) {
		want := "98 added to the queue"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		queue.Enqueue(98)

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Enqueue(98) = %q; want %q", got, want)
		}
	})

	t.Run("Test 6: Dequeue() removed to element 44 dequeue", func(t *testing.T) {
		want := "44 removed from the queue"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		queue.Dequeue()

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Dequeue() = %q; want %q", got, want)
		}
	})

	testsEnqueue := []struct {
		num  int
		want string
	}{
		{70, "70 added to the queue"},
		{22, "22 added to the queue"},
	}
	for _, tc := range testsEnqueue {
		t.Run(fmt.Sprintf("Test 7 Enqueue(%v) Add to the queue", tc.num), func(t *testing.T) {
			testStdout, writer, osStdout := utils.NewStDoutRestore(t)

			defer func() {
				os.Stdout = osStdout
			}()

			queue.Enqueue(tc.num)

			got := utils.Stdout(writer, testStdout)

			if got != tc.want {
				t.Fatalf("Enqueue(%q) = %q; want %q", tc.num, got, tc.want)
			}
		})
	}

	t.Run("Test 8: Dequeue() removed to element 32 dequeue", func(t *testing.T) {
		want := "32 removed from the queue"
		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		queue.Dequeue()

		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("Dequeue() = %q; want %q", got, want)
		}
	})

	testsEnqueue = []struct {
		num  int
		want string
	}{
		{67, "67 added to the queue"},
		{23, "23 added to the queue"},
	}
	for _, tc := range testsEnqueue {
		t.Run(fmt.Sprintf("Test 9: Enqueue(%v) Add to the queue", tc.num), func(t *testing.T) {
			testStdout, writer, osStdout := utils.NewStDoutRestore(t)

			defer func() {
				os.Stdout = osStdout
			}()

			queue.Enqueue(tc.num)

			got := utils.Stdout(writer, testStdout)

			if got != tc.want {
				t.Fatalf("Enqueue(%q) = %q; want %q", tc.num, got, tc.want)
			}
		})
	}
}
