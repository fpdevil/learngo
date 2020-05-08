package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Constant Length
//
//  Calculate how many characters inside the `home`
//  constant and print it.
//
// STEPS:
//  1. Declare a constant named `home`
//  2. Initialize it to "Milky Way Galaxy" string literal
//
//  3. Declare another constant named `length`
//  4. Initialize it by using the built-in function `len`.
//
//  5. Print the message below using the constants that
//     you've declared.
//
// RESTRICTION:
//  Use Printf.
//  Print the `home` constant using the quoted-string verb.
//
// EXPECTED OUTPUT
//  There are 16 characters inside "Milky Way Galaxy"
// ---------------------------------------------------------

// Exercise3 handles Constant Length
func Exercise3() {
	fmt.Println("")
	fmt.Println("---- 10-constants/03-constant-length Exercise: Constant Length ----")

	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)
	fmt.Printf("string verb: %q\n", home)
	fmt.Printf("There are %d characters inside %q\n", length, home)
}
