package main

import "fmt"

const arraySize = 10

//Hash Table Structure. Array of Buckets of Nodes
type hashTable struct {
	array [arraySize]*bucket
}

//Bucket Structure. Linked List of nodes
type bucket struct {
	head *node
}

//Node Structure. node of a linked list
type node struct {
	value string
	next  *node
}

//HASH TABLE METHODS
// insert
func (table *hashTable) insertTable(value string) {
	index := hash(value)
	table.array[index].insertToBucket(value)
}

// search
func (table *hashTable) searchTable(value string) {
	index := hash(value)
	table.array[index].searchBucket(value)
}

// delete

//BUCKET METHODS
// insert
func (nodeBucket *bucket) insertToBucket(value string) {
	if !(nodeBucket.searchBucket(value)) {
		newNode := &node{value: value}
		newNode.next = nodeBucket.head
		nodeBucket.head = newNode
	} else {
		fmt.Printf("%v is already in the list", value)
	}
}

// search
func (nodeBucket *bucket) searchBucket(value string) bool {
	currentNode := nodeBucket.head
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}

		currentNode = currentNode.next
	}
	return false
}

// delete
//printme

func (myList bucket) listNodes() {
	printMe := myList.head
	test := 10
	for test > 0 {
		fmt.Printf("%v\n", printMe.value)
		printMe = printMe.next
		test--
		if printMe == nil {
			break
		}
	}
}

// HASH FUNCTION
func hash(value string) int {
	total := 0
	for _, letter := range value {
		total += int(letter)
	}

	return total % arraySize
}

// INITIALIZE THE HASH TABLE
func initTable() *hashTable {
	ourTable := &hashTable{}
	for i := range ourTable.array {
		ourTable.array[i] = &bucket{}
	}

	return ourTable

}

func main() {
	ourTable := initTable()

	names := []string{"David", "PPPPP", "Sarah", "Tim", "Nebuchadnezzar", "jo", "dddddddd", "Ruth", "Lydia"}

	for _, name := range names {
		ourTable.insertTable(name)
	}

	ourTable.array[0].listNodes()

	//fmt.Printf("%+v\n", ourTable.array[7])
}
