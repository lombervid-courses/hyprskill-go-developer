package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func inputWithNewReader() {
	reader := bufio.NewReader(os.Stdin)

	b, err := reader.ReadBytes('\n') // Input into `b`: Hello World!\n
	if err != nil {
		log.Fatal(err) // Exit if we have an unexpected error
	}
	fmt.Println(string(b)) // Output: Hello World!\n

	s, err := reader.ReadString('d') // Input into `s`: JetBrains Academy\n
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s) // Output: JetBrains Acad
}

func inputWithNewScanner() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text() // Input: Sheldon Cooper 100 98 Physics\n
		fmt.Println(line)      // Output: Sheldon Cooper 100 98 Physics
	}
}

func ScanSpaceDelimitedWords() {
	wordScanner := bufio.NewScanner(os.Stdin)
	// Set the `Split` function to scan for words (space-delimited tokens):
	wordScanner.Split(bufio.ScanWords)

	for wordScanner.Scan() { // Input: Among Us ඞ\n
		fmt.Println(wordScanner.Text())
	}

	// Output:
	// Among
	// Us
	// ඞ
}

// The custom `ScanBools` function validates `bool` type input only:
/* func ScanBools(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseBool(string(token))
	}
	return advance, token, err
} */
func ScanWithCustomSplitFunction() {
	// The custom `ScanBools` function validates `bool` type input only:
	ScanBools := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseBool(string(token))
		}
		return advance, token, err
	}

	scanner := bufio.NewScanner(os.Stdin)
	// Set `ScanBools` as the split function for the scanning operation
	scanner.Split(ScanBools)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Exit if the scanned value is not a `bool`
	}

	// Input: true false Hello World!
	// Output:
	// true
	// false
	// 2022/02/24 23:02:04 strconv.ParseBool: parsing "Hello": invalid syntax
}

func main() {
	fmt.Println("Advanced Input")

	// inputWithNewReader()
	// inputWithNewScanner()
	// ScanSpaceDelimitedWords()
	ScanWithCustomSplitFunction()
}
