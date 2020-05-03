package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Fix the Float
//
//  Fix the program to print 2.5 instead of 2
//
// EXPECTED OUTPUT
//  2.5
// ---------------------------------------------------------

// Exercise2 from the 01-do-some-calculations
func Exercise2() {
	fmt.Println()
	fmt.Println("--- EXERCISE: Fix the Float ---")

	x := 5 / 2
	y := 5. / 2
	fmt.Printf("%d => %f", x, y)
}
