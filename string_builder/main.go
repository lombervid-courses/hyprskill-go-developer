package main

import (
	"fmt"
	"strings"
)

// func concat(strs ...string) string {
// 	var result string
// 	for _, str := range strs {
// 		result += str
// 	}
// 	return result
// }

func concat(strs ...string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}

func concatRunes(runes ...rune) string {
	var b strings.Builder
	for _, r := range runes {
		b.WriteRune(r)
	}
	return b.String()
}

func preallocateSize() {
	var b strings.Builder
	b.Grow(61) // We will be writing 61 bytes

	b.WriteString("Countdown to liftoff!\n") // 22 bytes written (including '\n')
	for i := 5; i >= 1; i-- {
		// 5 bytes written for each line (including '\n')
		b.WriteString(fmt.Sprintf("%d...\n", i))
	}
	b.WriteString("Liftoff! 🚀\n") // 14 bytes written (including '\n')

	fmt.Print(b.String())
	fmt.Println("Capacity of 'b' =", b.Cap())
	fmt.Println("Length of 'b' =", b.Len())
}

func main() {
	fmt.Println("String Builder")

	fmt.Println()
	fmt.Println(concat("Hello", " World", "!"))

	fmt.Println()
	fmt.Println(concatRunes('e', 'm', 'o', 'j', 'i', '😂', '👌', '💯')) // emoji😂👌💯

	fmt.Println()
	var b strings.Builder
	b.Write([]byte("Hello JetBrains Academy!"))
	fmt.Println(b.String()) // Hello JetBrains Academy!

	fmt.Println()
	preallocateSize()
}
