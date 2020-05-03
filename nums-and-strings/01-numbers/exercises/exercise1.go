package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Do Some Calculations
//
//  1. Print the sum of 50 and 25
//  2. Print the difference of 50 and 15.5
//  3. Print the product of 50 and 0.5
//  4. Print the quotient of 50 and 0.5
//  5. Print the remainder of 25 and 3
//  6. Print the negation of `5 + 2`
//
// EXPECTED OUTPUT
//  75
//  34.5
//  25
//  100
//  1
//  -7
// ---------------------------------------------------------

// Exercise1 from the 01-do-some-calculations
func Exercise1() {
	fmt.Println()
	fmt.Println("--- EXERCISE: Do Some Calculations ---")

	sum := 50 + 25
	fmt.Printf("sum of %d and %d is %d\n", 50, 25, sum)

	diff := 50 - 15.5
	fmt.Printf("difference of %d and %f is %g\n", 50, 15.5, diff)

	prod := int(50 * 0.5)
	fmt.Printf("product of %d and %f is %d\n", 50, 0.5, prod)

	quotient := int(50 / 0.5)
	fmt.Printf("quotient of %d and %f is %d\n", 50, 0.5, quotient)

	rem := 25 % 3
	fmt.Printf("remainder of %d and %d is %d\n", 25, 3, rem)

	neg := -(5 + 2)
	fmt.Printf("negation of '%d + %d' is %d", 5, 2, neg)
}
