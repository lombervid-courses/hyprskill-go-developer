package main

import (
	"fmt"
	"regexp"
)

var files = []string{
	"test.txt",
	"passwords.json",
	"notes.doc",
	"test2.txt",
	"dont't forget!!!.txt",
	"app.cfg",
	"send.rtf",
}

func matchThePattern() {
	fmt.Println("\nMatch the pattern")
	re := regexp.MustCompile(`go`)

	fmt.Println(re.Match([]byte("Golang")))
	fmt.Println(re.Match([]byte("golang")))

	fmt.Println(re.MatchString("It has returned false"))
	fmt.Println(re.MatchString("You've got it right!"))
}

func theDot() {
	fmt.Println("\nThe dot")
	re := regexp.MustCompile(`d..t`)

	for _, f := range files {
		if re.MatchString(f) {
			fmt.Println(f)
		}
	}

	// Output:
	// dont't forget!!!.txt
	// send.rtf
}

func theStar() {
	fmt.Println("\nThe star")
	re := regexp.MustCompile(`forget!*`)

	for _, f := range files {
		if re.MatchString(f) {
			fmt.Println(f)
		}
	}

	// Output:
	// dont't forget!!!.txt
}

func combinationOfMarks() {
	fmt.Println("\nCombination of marks")
	re := regexp.MustCompile(`t.*t`)

	for _, f := range files {
		if re.MatchString(f) {
			fmt.Println(f)
		}
	}

	// Output:
	// test.txt
	// test2.txt
	// dont't forget!!!.txt
}

func theBackslash() {
	fmt.Println("\nThe backslash")
	re := regexp.MustCompile(`.*\.json`)

	for _, f := range files {
		if re.MatchString(f) {
			fmt.Println(f)
		}
	}

	// Output:
	// passwords.json
}

func searchDIVTag() {
	fmt.Println("\nSearch DIV tag")
	re := regexp.MustCompile(`<div.*>.*<\/div.*>`)

	fmt.Println(re.Match([]byte("<div>Accept!</div>")))
	fmt.Println(re.Match([]byte("<div>Reject!</div>")))

	fmt.Println(re.MatchString("It's return false"))
	fmt.Println(re.MatchString("<div hidden class='h1'>With attributes!</div>"))

	// Output:
	// true
	// false
	// false
	// true
}

func main() {
	fmt.Println("Introduction to Regexp package")

	// compileRe, err := regexp.Compile(`**`)
	// if err != nil {
	// 	fmt.Println("Compile error:", err)
	// } else {
	// 	fmt.Println(compileRe)
	// }

	matchThePattern()
	theDot()
	theStar()
	combinationOfMarks()
	theBackslash()
	searchDIVTag()
}
