package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Operations with strings")

	fmt.Println("\nChanging the case of a string")

	question := "What Does The Fox Say?"
	fmt.Println(strings.ToLower(question))

	// Output:
	// what does the fox say?

	quote := "Are you not entertained?!\nIs this not why you are here?!"
	fmt.Println(strings.ToUpper(quote))

	// Output:
	// ARE YOU NOT ENTERTAINED?!
	// IS THIS NOT WHY YOU ARE HERE?!

	fmt.Println("\nTrimming strings")

	secret := "STATUS%1337%ENDSTATUS"
	fmt.Println(strings.Trim(secret, "ADENSTU%"))

	// Output:
	// 1337

	outerSpace := "    \tOuter Space...  \r\n"
	fmt.Println(strings.TrimSpace(outerSpace))

	// Output:
	// Outer Space...

	fmt.Println("\nSeparating and concatenating")

	phoneNumber := "8800-555-35-35"
	fmt.Println(strings.Split(phoneNumber, "-"))

	juiceDayDate := "18/September"
	fmt.Println(strings.Split(juiceDayDate, "/"))

	// Output:
	// [8800 555 35 35]
	// [18 September]

	taunt := []string{
		"Oh",
		"You want to fight",
		"Instead of running, you're coming towards me?!",
	}
	fmt.Println(strings.Join(taunt, "?!\n"))

	// Output:
	// Oh?!
	// You want to fight?!
	// Instead of running, you're coming towards me?!

	fmt.Println("\nWiping out words from strings")

	banned := "You Have Been Banned From The Minnie Mouse Club"
	fmt.Println(strings.Replace(banned, "Minnie", "Mickey", 1))

	// Output:
	// You Have Been Banned From The Mickey Mouse Club

	wrongChorus := `
Behind the world, around the world
Around the world, behind the world
Behind the world, around the world
`

	correctedChorus := strings.ReplaceAll(wrongChorus, "Behind", "Around")
	correctedChorus = strings.ReplaceAll(correctedChorus, "behind", "around")

	fmt.Println(correctedChorus)

	// Output:
	// Around the world, around the world
	// Around the world, around the world
	// Around the world, around the world
}
