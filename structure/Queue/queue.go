package queue

import "fmt"

// Queue
// Queue is abstract data type which depicts first in first out (FIFO) behavior.
//
//          Rear   First in Firts    Front
//           |  -------------------  |
//  Enqueue  8   6   5   4   3   2   1  ---> dequeue
//

var capacity int
var queueArr []int
var front int
var rear int
var currentSize = 0

func NewQueue(size int) {
	capacity = size
	front = 0
	rear = -1
	queueArr = make([]int, capacity)
}

// For adding element in the queue
func Enqueue(data int) {
	if isFull() {
		fmt.Print("Queue is full!! Can not add more elements")
		return
	}
	rear++
	if rear == capacity-1 {
		rear = 0
	}
	queueArr[rear] = data
	currentSize++
	fmt.Print(data, " added to the queue")
}

//This method removes an elements from the front of the queue
func Dequeue() {
	if isEmpty() {
		fmt.Print("Queue is empty!! Can not dequeue element")
		return
	}
	front++
	if front == capacity-1 {
		fmt.Print(queueArr[front-1], " removed from the queue")
		front = 0
	} else {
		fmt.Print(queueArr[front-1], " removed from the queue")
	}
	currentSize--
}

// For checking if queue is full
func isFull() bool {
	return currentSize == capacity
}
func isEmpty() bool {
	return currentSize == 0
}
