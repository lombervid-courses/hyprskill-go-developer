package main

import "fmt"

// Type alias
type anotherInt = int

// Custom type
type newInt int

func (n newInt) square() newInt {
	return n * n
}

func square(i int) int {
	return i * i
}

func main() {
	fmt.Println("Custom types")

	// var n = newInt(10)
	var n newInt = 23
	fmt.Println(n.square())
	fmt.Println(square(int(n)))

	var a anotherInt = 54
	fmt.Println(square(a))
}
