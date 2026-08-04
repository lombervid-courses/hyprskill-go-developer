package main

import (
	"fmt"
	"unicode"
)

func handleUnicodeSymbols() {

	unicodeString := "µ¶"

	fmt.Println("Regular iterator:")
	for i := 0; i < len(unicodeString); i++ {
		fmt.Printf("%d: %q\n", i, unicodeString[i])
	}

	fmt.Println("Range iterator:")
	for i, char := range unicodeString {
		fmt.Printf("%d: %q\n", i, char)
	}

	// Output:
	// Regular iterator:
	// 0: 'Â'
	// 1: 'µ'
	// 2: 'Â'
	// 3: '¶'
	// Range iterator:
	// 0: 'µ'
	// 2: '¶'
}

/*
 * `IsDigit` to check if the character is a digit (0-9);
 * `IsUpper` and IsLower to check the case of a character (uppercase: A-Z or lowercase:a-z);
 * `IsControl` to check if the character is a control symbol (\0\e\r);
 * `IsSpace` to check if the character is a whitespace symbol ( \n\t);
 * https://pkg.go.dev/unicode
 */
func symbolDefinition() {
	fmt.Println(unicode.IsDigit('1'))    // true
	fmt.Println(unicode.IsUpper('a'))    // false
	fmt.Println(unicode.IsLower('a'))    // true
	fmt.Println(unicode.IsControl('\n')) // true
	fmt.Println(unicode.IsSpace('\n'))   // true

	fmt.Println(unicode.In('ǈ', unicode.Latin)) // true
}

func symbolConvertion() {
	fmt.Println(string(unicode.ToLower('A'))) // a
	fmt.Println(string(unicode.ToUpper('a'))) // A
	fmt.Println(string(unicode.ToTitle('a'))) // A

	fmt.Println(string(unicode.ToLower('ǈ'))) // ǉ
	fmt.Println(string(unicode.ToUpper('ǈ'))) // Ǉ
	fmt.Println(string(unicode.ToTitle('ǈ'))) // ǈ
}

func main() {
	fmt.Println("Unicode package")
	// handleUnicodeSymbols()
	// symbolDefinition()
	symbolConvertion()
}
