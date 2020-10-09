package main

import "fmt"

type Stack struct {
	items []int
}

// Push data into Stack
func (ourStack *Stack) pushIntoStack(num int) {
	ourStack.items = append(ourStack.items, num)
}

// Pop data off the stack LIFO
func (ourStack *Stack) popOutOfStack() int {
	sizeOfSlice := len(ourStack.items) - 1
	removedItem := ourStack.items[sizeOfSlice]
	ourStack.items = ourStack.items[:sizeOfSlice]
	return removedItem
}

func main() {
	newStack := Stack{}

	newStack.pushIntoStack(5)
	newStack.pushIntoStack(78)
	newStack.pushIntoStack(61)
	newStack.pushIntoStack(2345)

	fmt.Printf("%+v\n", newStack)

	fmt.Println(newStack.popOutOfStack())
	fmt.Printf("%+v\n", newStack)
}
