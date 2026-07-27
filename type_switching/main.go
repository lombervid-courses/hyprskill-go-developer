package main

import "fmt"

func typeAssertion() {
	var v interface{} = "Golang"
	s, ok := v.(string)

	fmt.Println(v)
	fmt.Println(s, ok)

	i, ok := v.(int)
	fmt.Println(i, ok)

	// c := v.(int)
	// fmt.Println(c)
}

func typeSwitching() {
	var v interface{} = "Golang"
	// var v interface{} = 34
	switch v.(type) {
	case nil:
		fmt.Println("v is nil type")
	case int:
		fmt.Println("v is int type")
	case string:
		fmt.Println("v is string type")
	default:
		fmt.Println("type is not defined")
	}
}

func main() {
	fmt.Println("Type switching and type casting")
	// typeAssertion()
	typeSwitching()
}
