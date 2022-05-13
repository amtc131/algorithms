package stack

import "fmt"

// stack is abstract data type which depicts last in first out (LIFO) behavior
//                  Push   Pop
//                     \  /
//                    | 6  <------- top
//                    | 5
//  las in First OUT  | 4
//       (LIFO)       | 3
//                    | 2
//                   ▽ 1
//                   Stack

var size int
var arr []int
var top int

func MyStack(s int) {
	size = s
	arr = make([]int, size)
	top = -1
}

func Push(element int) {
	if !isFull() {
		top++
		arr[top] = element
		fmt.Print("Pushed element:", element)
	} else {
		fmt.Print("Stack is full")
	}
}

func Pop() int {
	if !isEmpty() {
		topElement := top
		top--
		fmt.Print("Popped element :", arr[topElement])
		return arr[topElement]
	} else {
		fmt.Print("Stack is empty!")
		return -1
	}
}

func Peek() int {
	if !isEmpty() {
		return arr[top]
	} else {
		fmt.Print("Stack is empty!")
		return -1
	}
}

func isFull() bool {
	return (size-1 == top)
}

func isEmpty() bool {
	return (top == -1)
}
