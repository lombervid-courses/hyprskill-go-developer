package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

const chunkSize = 12

func Reader() {
	reader := strings.NewReader("Supercalifragilisticexpialidocious")
	p := make([]byte, chunkSize)
	for chunkCount := 1; ; chunkCount++ {
		n, err := reader.Read(p)
		switch err {
		case nil:
			fmt.Println(chunkCount, " chunk: ", string(p[:n]))
		case io.EOF:
			fmt.Println("end of file")
			return
		default:
			fmt.Println(err)
			return
		}
	}

	// Output:
	// 1  chunk:  Supercalifra
	// 2  chunk:  gilisticexpi
	// 3  chunk:  alidocious
	// end of file

}

func Writer() {
	var b bytes.Buffer

	fmt.Fprintln(&b, "Winnie the Pooh: 'A hug is always the right size'.")
	fmt.Println(b.String())

	b.Write([]byte("Mary Poppins: 'Just a spoonful of sugar helps the medicine go down'."))
	fmt.Println(b.String())

	// Output:
	// Winnie the Pooh: 'A hug is always the right size'.
	//
	// Winnie the Pooh: 'A hug is always the right size'.
	// Mary Poppins: 'Just a spoonful of sugar helps the medicine go down'.
}

func Seeker() {
	quote := "Someplace where there isn't any trouble."
	offset := 16
	reader := strings.NewReader(quote)

	reader.Seek(int64(offset), io.SeekStart)

	b := make([]byte, len(quote)-offset-1)
	reader.Read(b)
	fmt.Println(string(b))

	// Output:
	// there isn't any trouble
}

func main() {
	// Reader()
	// Writer()
	Seeker()
}
