package main

import (
	"fmt"
	"log"
	"os"
)

func writeStringToFile() {
	s := "Hello, JB Academy!"
	if err := os.WriteFile("test.txt", []byte(s), 0644); err != nil {
		log.Fatal(err)
	}
}

func createAndWriteToFile() {
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := fmt.Fprint(file, "Hello, JB Academy!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes written successfully\n", b)
}

func writeLineByLine() {
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []string{"Hello, JB Academy!",
		"I am writing strings line by line",
		"I can write emojis too 🅱️®️🅾️!😤😤",
	}

	for i, line := range data {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			log.Fatal(err)
		}

		if i == len(data)-1 {
			fmt.Printf("%d lines written successfully\n", i+1)
		}
	}
}

func appendDataToAFile() {
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	additionalLine := "ALWAYS 🕓 make sure 👍 your code 💻 is 💡 clean 🧼 and well 💯 structured 🏛."
	b, err := fmt.Fprintln(file, additionalLine)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes appended successfully", b)
}

func main() {
	fmt.Println("Writing data to files")
	// writeStringToFile()
	// createAndWriteToFile()
	// writeLineByLine()
	appendDataToAFile()
}
