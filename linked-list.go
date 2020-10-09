package main

import "fmt"

// Implement the structure of a node data type
type node struct {
	data        string
	nextElement *node
}

// Implement the structure of the list
type linkedList struct {
	head   *node
	length int
}

// Add items to list
func (myList *linkedList) addNodeToList(n *node) {
	// Create pointer from the current new head node to the next item in list
	n.nextElement = myList.head
	// Setup the new head in our list
	myList.head = n
	//update the length of this list
	myList.length++
}

// delete items from list
func (myList *linkedList) deleteNodeFromList(node int) {
	currentHead := myList.head
	for i := 1; i < node-1; i++ {
		currentHead = currentHead.nextElement
	}
	currentHead.nextElement = currentHead.nextElement.nextElement
	myList.length--
}

// see whats in the list
func (myList linkedList) listNodes() {
	printMe := myList.head
	for myList.length > 0 {
		fmt.Printf("%v\n", printMe.data)
		printMe = printMe.nextElement
		myList.length--
	}
}

//func populateList(data []string) {
//	for i := 0; i < len(data); i++ {
//		ourList.addnodetolist(&node{data: data[i]})
//	}
//}

func main() {
	ourList := linkedList{}
	//	dataList := []string{"one", "two", "three"}
	ourList.addNodeToList(&node{data: "one"})
	ourList.addNodeToList(&node{data: "two"})
	ourList.addNodeToList(&node{data: "three"})
	ourList.addNodeToList(&node{data: "four"})

	ourList.deleteNodeFromList(2)
	//populateList(dataList)

	fmt.Printf("%+v\n", ourList)

	ourList.listNodes()
}
