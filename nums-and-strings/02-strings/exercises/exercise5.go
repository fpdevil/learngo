package exercises

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------

// Exercise5 handles Windows Path
func Exercise5(name string) {
	fmt.Println()
	fmt.Println("---- EXERCISE5: Improved Banger from 05-improved-banger ----")

	// msg := os.Args[1]

	// s := msg + strings.Repeat("!", len(msg))

	// fmt.Println(s)

	msg := name
	l := utf8.RuneCountInString(msg)
	s := msg + strings.Repeat("!", l)
	fmt.Println(s)
}
