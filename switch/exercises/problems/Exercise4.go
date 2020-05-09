package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

const usage4 = `
	[command] [string]

	Available commands: lower, upper and title
`

// Exercise4 - String Manipulator
func Exercise4(instr string) {
	fmt.Println("")
	fmt.Println("---- 12-switch/04-string-manipulator Exercise: String Manipulator ----")

	args := strings.Split(instr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if instr == "" || len(args) != 2 {
		fmt.Println(usage4)
		return
	}

	cmd, str := args[0], args[1]
	var result string

	switch cmd {
	default:
		fmt.Printf("Unknown command: %q\n", cmd)
	case "lower":
		result = strings.ToLower(str)
	case "upper":
		result = strings.ToUpper(str)
	case "title":
		result = strings.Title(str)

	}

	fmt.Printf("%s\n", result)
}
