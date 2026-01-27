package main

import (
	"fmt"

	"github.com/BestBor/datastructures-in-go/linkedlist"
)

func main() {

	// LinkedList playtest
	myList := linkedlist.New[int]()
	myList.Prepend(50)
	myList.Prepend(40)
	myList.Prepend(30)
	myList.Prepend(20)
	myList.Prepend(10)
	fmt.Println(myList.String())

}
