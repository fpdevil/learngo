package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Incdecs
//
//  1. Increase the `counter` 5 times
//  2. Decrease the `factor` 2 times
//  3. Print the product of counter and factor
//
// RESTRICTION
//  Use only the incdec statements
//
// EXPECTED OUTPUT
//  -75
// ---------------------------------------------------------

// Exercise4 from the 04-incdecs
func Exercise4() {
	fmt.Println()
	fmt.Println("--- EXERCISE: Incdecs ---")

	// DO NOT TOUCH THIS
	counter, factor := 45, 0.5

	// TYPE YOUR CODE BELOW
	// ...
	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	product := int(float64(counter) * factor)
	fmt.Printf("product of %d and %f is %d", counter, factor, product)

	// LASTLY: REMOVE THE CODE BELOW
	// _, _ = counter, factor
}
