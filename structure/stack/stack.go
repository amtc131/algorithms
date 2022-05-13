package stack

var size int
var arr []int
var top int

func MyStack(size int) {
	size = size
	arr = make([]int, size)
	top = -1
}
