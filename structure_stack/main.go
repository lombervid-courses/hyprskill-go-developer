package main

import "fmt"

func main() {
	fmt.Println("Using stack in Go")

	fmt.Println("\nStack")
	var myStack Stack

	myStack.Push(1)            // [1]
	myStack.Push(2)            // [1 2]
	popped, _ := myStack.Pop() // [1]
	myStack.Push(3)            // [1 3]

	fmt.Println(popped)        // 2
	fmt.Println(myStack.Pop()) // [1]   3 <nil>
	fmt.Println(myStack.Pop()) // []    1 <nil>

	fmt.Println("\nQueue")
	var myQueue Queue

	myQueue.Push(1)             // [1]   []
	myQueue.Push(2)             // [1 2] []
	poppedQ, _ := myQueue.Pop() // []    [2]
	myQueue.Push(3)             // [3]   [2]
	myQueue.Push(4)             // [3 4] [2]

	fmt.Println(poppedQ)       //            1
	fmt.Println(myQueue.Pop()) // [3 4] []   2 <nil>
	fmt.Println(myQueue.Pop()) // []    [4]  3 <nil>
	fmt.Println(myQueue.Pop()) // []    []   4 <nil>
}
