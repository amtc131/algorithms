package stack

import "fmt"

var size int
var arr []int
var top int

func MyStack(size int) {
	size = size
	arr = make([]int, size)
	top = -1
}

func Push(element int) {
	if !isFull() {
		top++
		arr[top] = element
		fmt.Printf("Pushed element: %d", element)
	} else {
		fmt.Println("Stack is full")
	}
}

func Pop() int {
	if !isEmpty() {
		topElement := top
		top--
		fmt.Printf("Popped element %d", arr[topElement])
		return arr[topElement]
	} else {
		fmt.Println("Stack is empty!")
		return -1
	}
}

func Peek() int {
	if !isEmpty() {
		return arr[top]
	} else {
		fmt.Println("Stack is empty!")
		return -1
	}
}

func isFull() bool {
	return (size-1 == top)
}

func isEmpty() bool {
	return (top == -1)
}
