package exercises

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

// Exercise4 handles Windows Path
func Exercise4(name string) {
	fmt.Println()
	fmt.Println("---- EXERCISE4: Count the Chars from 04-count-the-chars ----")

	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	// length := len(os.Args[1])
	// fmt.Println(length)

	length := utf8.RuneCountInString(name)
	fmt.Println(length)
}
