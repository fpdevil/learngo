package lessons

import "fmt"

// ---------------------------------------------------------
// EXERCISE: The Type Problem
//
//  Solve the data type problem in the program.
//
// EXPECTED OUTPUT
//  width: 265 height: 265
//  are they equal? true
// ---------------------------------------------------------

// Exercise2 The Type Problem
func Exercise2() {
	fmt.Println()
	fmt.Println("---- 02-the-type-problem Exercise: The Type Problem ----")

	// FIX THIS:
	// Change the following data types to the correct
	// data types where appropriate.
	var (
		width  uint16
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	// UNCOMMENT THIS:
	fmt.Println("are they equal?", width == height)
}
