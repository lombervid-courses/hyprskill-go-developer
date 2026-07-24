package main

import "fmt"

//func main() {
//	defer fmt.Println("Will be printed anyway!")
//	panic("Something has gone wrong!")
//	fmt.Println("Not printed because of the panic")
//}

// Output:
// Will be printed anyway!
// panic: Something has gone wrong!

/********************************** Recover **********************************/

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	defer func() {
		onPanic := recover() // catch the panic
		//fmt.Println(onPanic)
		if onPanic != nil {
			fmt.Printf("%d and %d are unacceptable for integer division\n", num1, num2)
		}
	}()
	//var num1, num2 int
	//fmt.Scan(&num1, &num2)

	fmt.Println(num1 / num2)
}

// Input:
// 1 0
// Output:
// 1 and 0 are unacceptable for integer division
