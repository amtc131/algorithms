package test

import (
	"os"
	"testing"

	module "github.com/amtc131/algorithms/module"
	"github.com/amtc131/algorithms/structure/utils"
)

func TestPrintNumbers(t *testing.T) {
	t.Run("Test 1 PrintNumbers(10)  from 1 to 10", func(t *testing.T) {
		want := "1 2 3 4 5 6 7 8 9 10 "

		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		module.PrintNumbers(10)
		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("PrintNumbers(10) = %q; want %q", got, want)
		}
	})

	t.Run("Test 1 PrintNumbers(20)  from 1 to 10", func(t *testing.T) {
		want := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 "

		testStdout, writer, osStdout := utils.NewStDoutRestore(t)

		defer func() {
			os.Stdout = osStdout
		}()

		module.PrintNumbers(20)
		got := utils.Stdout(writer, testStdout)

		if got != want {
			t.Fatalf("PrintNumbers(20) = %q; want %q", got, want)
		}
	})
}
