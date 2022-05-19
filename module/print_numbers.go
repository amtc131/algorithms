package module

import "fmt"

func PrintNumbers(n int) {
	if n != 0 {
		PrintNumbers(n - 1)
		fmt.Printf("%d ", n)
	}
}
