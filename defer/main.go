package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("defer/test.txt", os.O_RDWR|os.O_CREATE, 0644)
	//file, err := os.Create("defer/test.txt")
	//file, err := os.Create("/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
