package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

// Exercise9 handles Iota Seasons
func Exercise9() {
	fmt.Println("")
	fmt.Println("---- 10-constants/09-iota-seasons Exercise: Iota Seasons ----")

	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	// const (
	// 	Winter = 12
	// 	Spring = 3
	// 	Summer = 6
	// 	Fall   = 9
	// )
	const (
		Spring = 3 * (iota + 1)
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
