package main

import "fmt"

type Queue struct {
	items []int
}

// Push data into Queue
func (ourQueue *Queue) enqueue(num int) {
	ourQueue.items = append(ourQueue.items, num)
}

// Remove data from the queue FIFO
func (ourQueue *Queue) dequeue() int {
	sizeOfSlice := len(ourQueue.items)
	removedItem := ourQueue.items[0]
	ourQueue.items = ourQueue.items[1:sizeOfSlice]
	return removedItem
}

func main() {
	newQueue := Queue{}

	newQueue.enqueue(5)
	newQueue.enqueue(78)
	newQueue.enqueue(61)
	newQueue.enqueue(2345)

	fmt.Printf("%+v\n", newQueue)

	fmt.Println(newQueue.dequeue())
	fmt.Printf("%+v\n", newQueue)
}
