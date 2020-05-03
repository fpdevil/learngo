package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Windows Path
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

// Exercise1 handles Windows Path
func Exercise1() {
	fmt.Println()
	fmt.Println("---- EXERCISE1: Windows Path from 01-windows-path ----")

	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"
	fmt.Println(path)
	fmt.Println("With raw string literals")
	str := `c:\program files\duper super\fun.txt
c:\program files\really\funny.jpg`
	fmt.Println(str)
}
