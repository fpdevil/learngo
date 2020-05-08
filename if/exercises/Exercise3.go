package exercises

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go i wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

// Exercise3 - Arg Count
func Exercise3(vars string) {
	fmt.Println("")
	fmt.Println("---- 11-if/03-arg-count Exercise: Arg Count ----")

	v := strings.Split(vars, ",")
	// fmt.Printf("%#v, %T, %d\n", v, v, len(v))
	if v[0] == "" {
		fmt.Println("Give me args")
	} else if len(v) == 1 {
		fmt.Printf("There is one: %q\n", v[0])
	} else if len(v) == 2 {
		fmt.Printf("There are two: %s %s\n", v[0], v[1])
	} else {
		fmt.Printf("There are %d arguments\n", len(v))
	}

	return
}
