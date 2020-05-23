package problems

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

// Exercise9 - Append #3 — Fix the problems
func Exercise9() {
	fmt.Println("")
	fmt.Println("---- 16-slices/09-append-3-fix Exercise: Append #3 — Fix the problems ----")

	// toppings := []int{"black olives", "green peppers"}
	toppings := []string{"black olives", "green peppers"}

	// var pizza [3]string
	// append(pizza, ...toppings)
	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(toppings, "onions")
	// toppings = append(pizza, "extra cheese")
	pizza = append(pizza, "extra cheese")

	fmt.Printf("pizza       : %s\n", pizza)
}
