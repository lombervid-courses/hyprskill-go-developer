package main

import (
	"fmt"
	"internal/calculator"
	"internal/calculator/nested"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello World")

	calculator.Demo()
	//internal.SomeInternal()
	nested.NestedCalculator()

	color.Cyan("Finish")
}
