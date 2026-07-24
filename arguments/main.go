package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println("hello arguments")
	//fmt.Println(os.Args)
	//fmt.Println(os.Args[0])
	//fmt.Println(os.Args[1:])

	// Check that our program takes exactly three arguments:
	//if len(os.Args) != 4 {
	//	log.Fatal("Error! Expected 3 arguments only!")
	//}
	//
	//fmt.Printf("Contents of the os.Args slice = %v\n", os.Args)
	//fmt.Printf("Name of our program --> os.Args[0] = %[1]s | type: %[1]T\n", os.Args[0])
	//fmt.Printf("First cmd-line argument --> os.Args[1] = %[1]s | type: %[1]T\n", os.Args[1])
	//fmt.Printf("Second cmd-line argument --> os.Args[2] = %[1]s | type: %[1]T\n", os.Args[2])
	//fmt.Printf("Third cmd-line argument --> os.Args[3] = %[1]s | type: %[1]T\n", os.Args[3])

	/** Flags **/

	// Declare string and int flags with default values and help messages:
	//name := flag.String("name", "User", "Enter your name")
	////age := flag.Int("age", 1, "Enter your age")
	//
	//var age int
	//flag.IntVar(&age, "age", 1, "Enter your age")
	//
	//// Another way to declare a flag - bind a command-line flag to an existing variable:
	//var spacing bool
	//flag.BoolVar(&spacing, "spacing", true, "Insert a new line between each message")
	//
	//// After declaring all the flags, enable command-line flag parsing:
	//flag.Parse()
	//
	//fmt.Printf("Hello, %s! ", *name)
	//if spacing {
	//	fmt.Println()
	//}
	//fmt.Printf("You are %d years old.", age)

	/** Subcommands **/

	// Declare the 'repeat' subcommand via the NewFlagSet() function:
	repeat := flag.NewFlagSet("repeat", flag.ExitOnError)

	// Declare two flags 'name' and 'count' on the 'repeat' subcommand:
	repeatName := repeat.String("name", "User", "Enter the name to be repeated")
	repeatCount := repeat.Int("count", 1, "Enter the number of repetitions")

	if os.Args[1] == "repeat" {
		repeat.Parse(os.Args[2:]) // Parse the arguments after the subcommand
		for i := 0; i < *repeatCount; i++ {
			fmt.Printf("%s\n", *repeatName)
		}
	} else {
		log.Fatal("Expected 'repeat' subcommand") // Exit if the subcommand is not 'repeat'
	}
}
