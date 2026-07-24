package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func wholeContent() {
	data, err := os.ReadFile("text_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func lineByLine() {
	file, err := os.Open("text_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scanner.Split(bufio.ScanWords) // read word by word

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

const chunkSize = 15

// const chunkSize = 45

func inChunks() {
	file, err := os.Open("text_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, chunkSize) // create a slice of bytes buffer with
	// the previously defined chunk size

	for {
		readTotal, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break // after reading the last chunk, break the loop
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:readTotal]))
		// fmt.Print(string(buf[:readTotal]))
	}

}

func main() {
	fmt.Println("Read file")
	// wholeContent()
	// lineByLine()
	inChunks()
}
