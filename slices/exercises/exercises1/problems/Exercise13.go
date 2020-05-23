package problems

import (
	"fmt"
	"strconv"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

// Exercise13 - Slice the numbers
func Exercise13() {
	fmt.Println("")
	fmt.Println("---- 16-slices/13-slicing-basics Exercise: Slice the numbers ----")

	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	var (
		nums, evens, odds, middle, first, last, evenlast, oddslast []int
	)

prime:
	for _, v := range strings.Split(data, " ") {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue prime
		}
		nums = append(nums, n)

		if n%2 == 0 {
			evens = append(evens, n)
		} else {
			odds = append(odds, n)
		}
	}

	// this has to be created after nums is populated
	length := len(nums)
	if length%2 == 0 {
		middle = append(middle, nums[length/2-1:length/2]...)
	} else {
		middle = append(middle, nums[length-1/2])
	}

	if length >= 2 {
		first = append(first, nums[:2]...)
	}

	if length >= 2 {
		last = append(last, nums[length-2:]...)
	}

	evenlen := len(evens)
	oddlen := len(odds)
	if evenlen >= 1 {
		evenlast = append(evenlast, evens[evenlen-1:]...)
	}
	if oddlen >= 2 {
		oddslast = append(oddslast, odds[evenlen-2:]...)
	}

	fmt.Printf("nums			:%v\n", nums)
	fmt.Printf("evens			:%v\n", evens)
	fmt.Printf("evens			:%v\n", odds)
	fmt.Printf("middle			:%v\n", middle)
	fmt.Printf("first 2			:%v\n", first)
	fmt.Printf("last 2			:%v\n", last)
	fmt.Printf("evens last 1	:%v\n", evenlast)
	fmt.Printf("odds last 2		:%v\n", oddslast)

	s.Show("nums", nums)
	s.Show("evens", evens)
	s.Show("evens", odds)
	s.Show("middle", middle)
	s.Show("first 2", first)
	s.Show("last 2", last)
	s.Show("evens last 1", evenlast)
	s.Show("odds last 2", oddslast)
}
