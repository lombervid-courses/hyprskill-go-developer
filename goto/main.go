package main

import "fmt"

func label() {
	fmt.Println("\nLabel syntax")
	var num, divCount, targetCount int
	fmt.Scan(&targetCount)

NumberLoop: //label
	for {
		divCount = 0
		for i := 2; i < num; i++ {
			if num%i == 0 {
				divCount++
			}
			if divCount >= targetCount {
				break NumberLoop // break by the label
				// break // break without the label
			}
		}

		num++
	}

	fmt.Println(num)

	// Input:
	// 123
	// Output:
	// 83160
}

func gotoSyntax() {
	fmt.Println("\nGoto syntax")

	fmt.Println("I'm printed")

	goto EndOfTheCode

	fmt.Println("I'm not printed")

EndOfTheCode:
	fmt.Println("Print at the end")

	// Output:
	// I'm printed
	// Print at the end
}

func gotoBestPractices() {
	fmt.Println("\nBest practices")
	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("is even")
		goto TheEnd
	}
	// else
	fmt.Println("is odd")

TheEnd:
	fmt.Println("end")

	// Input:
	// 2
	// Output:
	// is even
	// end

	// Input:
	// 3
	// Output:
	// is odd
	// end
}

func main() {
	// fmt.Println("\n")
	fmt.Println("Goto and labels")

	// label()
	// gotoSyntax()
	gotoBestPractices()
}
