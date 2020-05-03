package exercises

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: ToLowercase
//
//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
//
// HINT
//  Check out the strings package from Go online documentation.
//  You will find the lowercase function there.
//
// INPUT
//  "SHEPARD"
//
// EXPECTED OUTPUT
//  shepard
// ---------------------------------------------------------

// Exercise6 handles Windows Path
func Exercise6(name string) {
	fmt.Println()
	fmt.Println("---- EXERCISE6: ToLowercase from 06-tolowercase ----")

	s := strings.ToLower(name)
	fmt.Println(s)
}
