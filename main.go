package main

import (
	"fmt"

	"github.com/BestBor/datastructures-in-go/linkedlist"
	"github.com/BestBor/datastructures-in-go/stack"
)

func main() {

	// LinkedList playtest
	fmt.Println("LinkedList:")
	myList := linkedlist.New[int]()
	myList.Prepend(50)
	myList.Prepend(40)
	myList.Prepend(30)
	myList.Prepend(20)
	myList.Prepend(10)
	fmt.Println(myList.String())
	fmt.Println(myList)

	// Stack playtest
	fmt.Println("Stacks:")
	myStack := stack.New[int]()
	myStack.Push(50)
	myStack.Push(40)
	myStack.Push(30)
	myStack.Push(20)
	myStack.Push(10)
	fmt.Println(myStack.IsEmpty())
	fmt.Println(myStack.Len())
	fmt.Println(myStack.Peek())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)

}
