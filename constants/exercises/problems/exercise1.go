package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Minutes in Weeks
//
//  Calculate how many minutes in two weeks.
//
//  STEPS:
//  1. Declare `minsPerDay` constant and initialize it
//     to the number of minutes in a day
//
//  2. Declare `weekDays` constant and initialize it
//     to 7.
//
//  3. Use printf to print the total number of minutes
//     in two weeks.
//
// EXPECTED OUTPUT
//  There are 20160 minutes in 2 weeks.
// ---------------------------------------------------------

// Exercise1 handles Minutes in Weeks
func Exercise1() {
	fmt.Println("")
	fmt.Println("---- 10-constants/01-minutes-in-weeks Exercise: Minutes in Weeks ----")

	const minsPerDay = 60 * 24
	const weekDays = 7

	fmt.Printf("There are %d minutes in %d weeks\n", minsPerDay*2*weekDays, 2)
}
