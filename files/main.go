package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Files")

	// Open a file
	file, err := os.Open("test_file.txt")
	// file, err := os.OpenFile("test_file.txt", os.O_RDONLY, 0644)

	// Create a file
	// file, err := os.Create("test_file.txt")
	// file, err := os.OpenFile("test_file.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write to a file
	// file, err := os.Create("test_file.txt")
	// file, err := os.OpenFile("test_file.txt", os.O_RDWR, 0644)
	// file, err := os.OpenFile("test_file.txt", os.O_CREATE|os.O_RDWR, 0666)
	// file, err := os.OpenFile("test_file.txt", os.O_RDWR|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// fmt.Fprintln(file, "weeeeee")
	// fmt.Fprintln(file, "weeeeee 2")
	// fmt.Fprintln(file, "weeeeee 3")
	// fmt.Fprintln(file, "weeeeee 4")

	// Delete a file/directory
	// err := os.Remove("test_file.txt")
	// err := os.Remove("mydir")
	// err := os.RemoveAll("mydir")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
