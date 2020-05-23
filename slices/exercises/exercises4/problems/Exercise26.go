package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
//      reqs := []int{
// 	      500, 600, 250,
// 	      200, 400, 50,
// 	      900, 800, 600,
// 	      750, 250, 100,
// 	      150, 654, 235,
// 	      320, 534, 765,
// 	      121, 876, 285,
// 	      543, 642,
// 	      // the last element is missing (your code should be able to handle this)
// 	      // that is why you shouldn't calculate the `size` below manually.
//      }
//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

// Exercise26 - Print daily requests
func Exercise26() {
	fmt.Println("")
	fmt.Println("---- 16-slices/26-print-daily-requests Exercise: Print daily requests ----")

	// args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// ================================================
	// #1: Make a new slice with the exact size needed.

	// _ = reqs // remove this when you start

	size := 0 // you need to find the size.
	for i := range reqs {
		if i%N == 0 {
			size++
		}
	}
	// fmt.Println("size: ", size)

	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	_ = daily // remove this when you start

	l := len(reqs)
	for from := 0; from < l; from += N {
		to := from + N
		// fmt.Printf("[%d: %d] ", from, to)
		if to > l {
			to = l
		}

		current := reqs[from:to]
		// fmt.Printf("%v\n", current)
		daily = append(daily, current)
	}

	// fmt.Printf("daily: %v\n", daily)

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	// ...
	var grand int
	for i, v := range daily {
		sum := 0
		for j := range v {
			sum += v[j]
		}
		grand += sum
		fmt.Printf("%-10d%-10d\n", i, sum)
	}

	fmt.Println(strings.Repeat("=", 20))
	fmt.Printf("%-10s%-10d\n", "Grand", grand)
	// ================================================
}
