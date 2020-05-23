package problems

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Append and Sort Numbers
//
//  We'll have a []string that should contain numbers.
//
//  Your task is to convert the []string to an int slice.
//
//  1. Get the numbers from the command-line
//
//  2. Append them to an []int
//
//  3. Sort the numbers
//
//  4. Print them
//
//  5. Handle the error cases
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    provide a few numbers
//
//  go run main.go 4 6 1 3 0 9 2
//    [0 1 2 3 4 6 9]
//
//  go run main.go a b c
//    []
//
//  go run main.go 4 a 1 c d 9
//    [1 4 9]
//
// ---------------------------------------------------------

// Exercise10 - Append and Sort Numbers
func Exercise10(vals string) {
	fmt.Println("")
	fmt.Println("---- 16-slices/10-append-sort-nums Exercise: Append and Sort Numbers ----")

	args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// handle arguments
	if vals == "" || len(args) < 1 {
		fmt.Println("provide a few numbers")
		return
	}

	var nums []int

prime:
	for _, v := range args {
		num, err := strconv.Atoi(v)
		if err != nil {
			// fmt.Printf("unable to convert %d: %v\n", i, v)
			continue prime
		}
		nums = append(nums, num)
	}

	// sort the values from slice
	sort.Ints(nums)
	fmt.Printf("%v\n", nums)
}
