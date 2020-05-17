package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

// Exercise12 - Number Sorter
func Exercise12(vals string) {
	fmt.Println()
	fmt.Println("---- 14-arrays/12-sorter Exercise: Number Sorter	 ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	// handle exactly 5 arguments
	if len(args) > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	var (
		nums [5]float64
		nlen int
	)

	for i, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		} else {
			nums[i] = num
		}
	}

	// Bubble Sort
	// iterate over the array n-1 times
	// for each iteration, check adjacent pair and swap them
	// if the first element is greater than the second one
	nlen = len(nums)
	for range nums {
		for i, v := range nums {
			// index under the length and the value more than next
			if i < nlen-1 && v > nums[i+1] {
				// fmt.Printf("index: %d, is %f > %f\n", i, v, nums[i+1])
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	fmt.Printf("%v\n", nums)
}
