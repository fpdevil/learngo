package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	fmt.Println("--- 19-strings-runes-bytes/03-rune-manipulator ---")
	fmt.Println("    EXERCISE: Rune Manipulator    ")

	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	// _ = strings

	for i, s := range strings {
		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("Byte Length of %s 		: 	%d bytes\n", strings[i], len(strings[i]))
		fmt.Printf("Rune Length of %s 		: 	%d bytes\n", strings[i], utf8.RuneCountInString(strings[i]))
		fmt.Println()

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("Hexadecimal bytes of %s 	: 	% [1] x\n", s)
		fmt.Println()

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		// navigate for each string...
		fmt.Println("runes of the strings in hexadecimal")
		for _, r := range s {
			fmt.Printf("% x\n", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Println("runes of the strings as rune literals")
		for _, r := range s {
			fmt.Printf("%[1]q\n", r)
		}

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		fmt.Println("first rune and its byte size of the strings")
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("first rune	: 	%q and size	: 	%d bytes\n", r, size)
		fmt.Println()

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		fmt.Println("last rune of the strings")
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("last rune	: 	%q and size	: 	%d bytes\n", r, size)
		fmt.Println()

		// Slice and print the first two runes of the strings
		fmt.Println("first two runes of the strings")
		_, firstsize := utf8.DecodeRuneInString(s)
		_, secondsize := utf8.DecodeRuneInString(s[firstsize:])
		fmt.Printf("first two runes : %q\n", s[:firstsize+secondsize])
		fmt.Println()

		// Slice and print the last two runes of the strings
		fmt.Println("last two runes of the strings")
		_, last := utf8.DecodeLastRuneInString(s)
		_, lastbutone := utf8.DecodeLastRuneInString(s[:len(s)-last])
		fmt.Printf("last two runes : %q\n", s[len(s)-lastbutone-last:])
		fmt.Println()

		// Convert the string to []rune
		// Print the first and last two runes
		rn := []rune(s)
		fmt.Printf("first two runes: %q\n", string(rn[:2]))
		fmt.Printf("last two runes: %q\n", string(rn[len(rn)-2:]))
	}
}
