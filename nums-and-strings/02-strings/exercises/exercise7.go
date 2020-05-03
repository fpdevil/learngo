package exercises

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     the given string
//  3. Trim the text variable and print it
//
// EXPECTED OUTPUT
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------

// Exercise7 handles Trim It
func Exercise7() {
	fmt.Println()
	fmt.Println("---- EXERCISE7: Trim It from 07-trim-it ----")

	msg := `
	
	The weather looks good.
I should go and play.
	`

	// fmt.Println(msg)
	s := strings.TrimSpace(msg)
	fmt.Println(s)
}
