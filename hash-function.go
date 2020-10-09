package main

import (
	"fmt"
)

const sizeOfArray = 10

func hash(name string) int {
	//iterate over string and cast to integer
	sumOfLetters := 0
	for _, letter := range name {
		sumOfLetters += int(letter)
	}

	return sumOfLetters % sizeOfArray
}

func main() {
	fmt.Println(hash("Nebuchadnezzar"))
}
