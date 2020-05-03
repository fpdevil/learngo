package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Precedence
//
//  Change the expressions to produce the expected outputs
//
// RESTRICTION
//  Use parentheses to change the precedence
// ---------------------------------------------------------

// Exercise3 from the 03-precedence
func Exercise3() {
	fmt.Println()
	fmt.Println("--- EXERCISE: Precedence ---")

	// This expression should print 20
	// 10 + 5 - 5 - 10
	fmt.Println(10 + 5 - (5 - 10))

	// This expression should print -16
	// -10 + 0.5 - 1 + 5.5
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// This expression should print -25
	fmt.Println(5 + 10*(2-5))

	// This expression should print 0.5
	// 0.5*2 - 1
	fmt.Println(0.5 * (2 - 1))

	// This expression should print 24
	// 3 + 1/2*10 + 4
	fmt.Println((3+1)/2*10 + 4)

	// This expression should print 15
	// 10 / 2 * 10 % 7
	fmt.Println((10 / 2) * (10 % 7))

	// This expression should print 40
	// Note that you may need to use floats to solve this
	// 100 / 5 / 2
	fmt.Println(100 / (5 / 2.))
}
