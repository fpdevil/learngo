package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Months
//
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
//
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
//
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------

// Exercise7 handles Iota Months
func Exercise7() {
	fmt.Println("")
	fmt.Println("---- 10-constants/07-iota-months Exercise: Iota Months ----")

	// const (
	// 	Nov = 11
	// 	Oct = 10
	// 	Sep = 9
	// )
	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)
}
