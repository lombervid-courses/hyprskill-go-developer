package main

import (
	"fmt"
	"strings"
)

func prefixAndSufix() {
	var source string = "Hyperskill"
	var prefix string = "Hype"

	hasPrefix := strings.HasPrefix(source, prefix)
	hasSuffix := strings.HasSuffix(source, "LL")

	fmt.Println("\nPrefix and Suffix")
	fmt.Println(hasPrefix)
	fmt.Println(hasSuffix)

	// Output:
	// true
	// false
}

func contains() {
	var source string = "Test string for Contains"

	fmt.Println("\nContains")
	fmt.Println(strings.Contains(source, "for"))
	fmt.Println(strings.Contains(source, "test"))

	// Output:
	// true
	// false
}

func index() {
	var source string = "Who is who in Doctor Who?"

	fmt.Println("\nIndex")
	fmt.Println(strings.Index(source, "Who"))
	fmt.Println(strings.LastIndex(source, "Who"))
	fmt.Println(strings.Index(source, "doctor"))

	// Output:
	// 0
	// 21
	// -1
}

func allIndexes() {
	var source string = "Will Will Smith smith Will Smith?"

	i := 0
	n := strings.Index(source, "Will")

	fmt.Println("\nAll indexes")
	for i != -1 {
		fmt.Println(n)

		i = strings.Index(source[n+1:], "Will")
		n = n + i + 1
	}

	// Output:
	// 0
	// 5
	// 22
}

func count() {
	fmt.Println("\nCount")
	fmt.Println(strings.Count("Perpetuum mobile", "e"))
	fmt.Println(strings.Count("Perpetuum mobile", "z"))
	fmt.Println(strings.Count("Perpetuum mobile", ""))

	// Output:
	// 3
	// 0
	// 17
}

func caseInsensitiveComparison() {
	fmt.Println("\nCase insensitive comparison")

	fmt.Println(strings.EqualFold("Hello", "hello"))
	fmt.Println(strings.EqualFold("Hi!", "hi"))

	// Output:
	// true
	// false
}

func main() {
	fmt.Println("String search")

	prefixAndSufix()
	contains()
	index()
	allIndexes()
	count()
	caseInsensitiveComparison()
}
