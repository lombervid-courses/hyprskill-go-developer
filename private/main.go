package main

import (
	"fmt"
	"private/creatures"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {

	// create an instance of the `Animal` struct imported from the `creatures` package
	// this works because the `Animal` struct is public
	var crocodile creatures.Animal

	fmt.Println(crocodile)
	//crocodile.Human = creatures.human{"Jerry", 34}
	//fmt.Println(crocodile)

	// trying to create an instance of the `human` struct
	// this will fail because the `human` struct is private
	//var jerry creatures.human
	//fmt.Println(jerry)

	number := 0

	if true {
		number := 10
		number++
	}

	fmt.Println(number) // 0
}
