package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Simplify the Assignments
//
//  Simplify the code (refactor)
//
// RESTRICTION
//  Use only the incdec and assignment operations
//
// EXPECTED OUTPUT
//  3
// ---------------------------------------------------------

// Exercise6 from the 06-simplify-the-assignments
func Exercise6() {
	fmt.Println()
	fmt.Println("--- EXERCISE: Simplify the Assignments ---")

	width, height := 10, 2

	width++         // width = width + 1
	width += height // width = width + height
	width--         // width = width - 1
	width -= height // width = width - height
	width *= 20     // width = width * 20
	width /= 25     // width = width / 25
	width %= 5      // width = width % 5

	fmt.Println(width)
}
