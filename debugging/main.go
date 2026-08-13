package main

import (
	"fmt"
	"log"
	"os"
)

/*
 * Logging Packages:
 * - https://pkg.go.dev/log/slog
 * - https://github.com/uber-go/zap
 * - https://github.com/sirupsen/logrus
 * - https://github.com/rs/zerolog
 * - https://github.com/charmbracelet/log
 */
func logging() {
	fmt.Println("\nLogging as a debugging tool")

	isEven := func(num int) bool {
		logger := log.New(os.Stdout, "isEven: ", log.Lshortfile+log.Ltime+log.Ldate)
		res := num % 2
		if res == 1 {
			logger.Printf("num = %d, res = %d, return true\n", num, res)
			return true
		}
		logger.Printf("num = %d, res = %d, return false\n", num, res)
		return false
	}

	var array = []int{1, 2, 3, 4, 5, 6}
	for _, num := range array {
		if isEven(num) {
			fmt.Printf("number %d is even\n", num)
		} else {
			fmt.Printf("number %d is odd\n", num)
		}
	}
}

/*
 * Delve as a common Go debug tool
 *
 * - https:github.com/go-delve/delve
 *
 * `go install github.com/go-delve/delve/cmd/dlv@latest`
 */
func delve() {
	// dlv break main.delve:2
	fmt.Println("\nDelve as a common Go debug tool")

	// dlv break ./main.go:53
	fmt.Println("other")
}

func main() {
	// fmt.Println("\n")
	fmt.Println("Debugging Go code")

	logging()
	delve()
}
