package main

import (
	"fmt"
	"log"
	"strconv"
)

func parseIntegers() {
	num := "123" // 'num' contains a string of integer numbers

	fmt.Printf("Initial type: %T | value: %v\n", num, num)
	val, err := strconv.Atoi(num) // We save the converted value in the 'val' variable
	if err != nil {
		log.Fatal(err) // Exit if we have an error
	}
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)

	// Output:
	// Initial type: string | value: 123
	// Converted type: int  | value: 123
}

func parseIntegers2() {
	num := "123" // 'num' contains a string of integer numbers

	fmt.Printf("Initial type: %T   | value: %v\n", num, num)
	val, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)
	// Output:
	// Initial type: string   | value: 123
	// Converted type: int64  | value: 123
}

func parseFloats() {
	num := "3.1416" // 'num' contains a string of floating point numbers

	fmt.Printf("Initial type: %T | value: %v\n", num, num)
	val, err := strconv.ParseFloat(num, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)
	// Output:
	// Initial type: string | value: 3.1416
	// Converted type: float64  | value: 3.1416

	if val, err := strconv.ParseFloat("NaN", 32); err == nil {
		fmt.Printf("Type: %T | value: %v\n", val, val)
	}

	if val, err := strconv.ParseFloat("Infinity", 32); err == nil {
		fmt.Printf("Type: %T | value: %v\n", val, val)
	}
}

func intToStr() {
	num := 456 // 'num' contains an integer value

	fmt.Printf("Initial type: %T       | value: %v\n", num, num)
	val := strconv.Itoa(num) // Converts 'num' to a string and assigns it to 'val'
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)
	// Output:
	// Initial type: int       | value: 456
	// Converted type: string  | value: 456
}

func floatToStr() {
	num := 3.1415926535

	fmt.Printf("Initial type: %T   | value: %v\n", num, num)
	val := strconv.FormatFloat(num, 'f', -1, 32)
	fmt.Printf("Converted type: %T  | value: %v\n\n", val, val)

	fmt.Printf("Initial type: %T   | value: %v\n", num, num)
	val = strconv.FormatFloat(num, 'f', -1, 64)
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)

	// Output:
	// Initial type: float64   | value: 3.1415926535
	// Converted type: string  | value: 3.1415927

	// Initial type: float64   | value: 3.1415926535
	// Converted type: string  | value: 3.1415926535
}

func parseBool() {
	b := "false"

	fmt.Printf("Initial type: %T  | value: %v\n", b, b)
	val, err := strconv.ParseBool(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)

	// Output:
	// Initial type: string  | value: false
	// Converted type: bool  | value: false
}

func boolToStr() {
	b := false

	fmt.Printf("Initial type: %T      | value: %v\n", b, b)
	val := strconv.FormatBool(b)
	fmt.Printf("Converted type: %T  | value: %v\n", val, val)

	// Output:
	// Initial type: bool      | value: false
	// Converted type: string  | value: false
}

func main() {
	fmt.Println("Parsing data from strings")

	// parseIntegers()
	// parseIntegers2()
	// parseFloats()
	// intToStr()
	// floatToStr()
	// parseBool()
	boolToStr()
}
