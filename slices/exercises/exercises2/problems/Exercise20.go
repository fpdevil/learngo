package problems

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 10e6 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

// Exercise20 - Observe the capacity growth
func Exercise20() {
	fmt.Println("")
	fmt.Println("---- 16-slices/20-observe-the-cap-growth Exercise: Observe the capacity growth ----")

	// args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// 1. Create a nil slice
	var nums []int

	// define initial slice capacity of 1. to keep track
	initCap := 1.

	// 2. Loop 10 million or 10e6 times
	const times = 1e6

	// perform the loop
	for i := 0; i <= times; i++ {
		c := float64(cap(nums))

		// during resizing at some points capacity becomes 0 until
		// next resize by go vm
		if c != initCap || c == 0 {
			//  4. Print the length and capacity of the slice "only"
			//     when its capacity changes.
			fmt.Printf("len: %-15d capacity: %-15g growth: %.2f\n", len(nums), c, c/initCap)
		}

		// now initial capacity becomes current capacity
		initCap = c

		// 3. On each iteration: Append an element to the slice
		nums = append(nums, 1)
	}
}
