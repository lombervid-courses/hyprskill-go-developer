package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Math package")

	fmt.Println("\nAdvanced arithmetic functions")

	fmt.Println(math.Abs(-1.618)) // Abs is 1.618
	fmt.Println(math.Min(0, -6))  // Min is -6
	fmt.Println(math.Max(3, 2))   // Max is 3

	fmt.Println("\nRounding functions")

	fmt.Println(math.Floor(3.43)) // Floor is 3
	fmt.Println(math.Ceil(1.71))  // Ceil is 2

	roundUp := math.Round(7.50)
	fmt.Println(roundUp) // roundUp is 8

	roundDown := math.Round(7.49)
	fmt.Println(roundDown) // roundDown is 7

	fmt.Println("\nConstants in the math package")

	fmt.Println(math.Pi) // Pi is 3.141592653589793...
	fmt.Println(math.E)  // E is 2.718281828459045...

	fmt.Println(math.MinInt) // MinInt is -9223372036854775808
	fmt.Println(math.MaxInt) // MaxInt is 922337203685477580

	fmt.Println("\nExponential and logarithmic functions")

	fmt.Println(math.Sqrt(4))   // Sqrt is 2
	fmt.Println(math.Pow(2, 2)) // the square of 2 is 4

	fmt.Println(math.Cbrt(64))  // Cbrt is 4
	fmt.Println(math.Pow(4, 3)) // the cube of 4 is 64

	fmt.Println(math.Log(1))      // Log of 1 is 0
	fmt.Println(math.Log(math.E)) // Log of math.E is 1

	fmt.Println("\nTrigonometric functions")

	degrees := 60.0

	// Below we apply the degrees to radians formula:
	radians := degrees * math.Pi / 180 // 60.0 degrees are 1.04719... radians

	fmt.Println(math.Cos(radians)) // Cos is 0.5000...
	fmt.Println(math.Sin(radians)) // Sin is 0.8660...
	fmt.Println(math.Tan(radians)) // Tan is 1.7320...
}
