package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Months #2
//
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------

// Exercise8 handles Iota Months #2
func Exercise8() {
	fmt.Println("")
	fmt.Println("---- 10-constants/08-iota-months-2 Exercise: Iota Months #2 ----")

	// HINT: This is a valid constant declaration
	//       Blank-Identifier can be used in place of a name
	// const _ = iota
	//    ^- this is just a name

	// Now, use iota and initialize the following constants
	// "automatically" to 1, 2, and 3 respectively.
	// const (
	// 	Jan = iota
	// 	Feb
	// 	Mar
	// )
	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}
