package problems

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Fix the problems
//
//  1. Uncomment the code
//
//  2. Fix the problems
//
//  3. BONUS: Simplify the code
//
//
// EXPECTED OUTPUT
//   "Einstein and Shepard and Tesla"
//   ["Fire" "Kafka's Revenge" "Stay Golden"]
//   [1 2 3 5 6 7 8 9]
// ---------------------------------------------------------

// Exercise5 - Fix the problems
func Exercise5() {
	fmt.Println("")
	fmt.Println("---- 16-slices/05-fix-the-problems Exercise: Fix the problems ----")

	//-- the following code was uncommented {start}
	var names []string
	// names := []string{}
	names = []string{
		"Einstein", "Shepard",
		"Tesla",
	}

	// -----------------------------------
	var books = []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	// -----------------------------------
	// this time, do not change the nums array to a slice
	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	// use the slicing expression to change the nums array to a slice below
	// sort.Ints(nums)
	sort.Ints(nums[:]) // converting to slice with [:]

	// -----------------------------------
	// Here: Use the strings.Join function to join the names
	//       (see the expected output)
	fmt.Printf("%q\n", names)
	fmt.Printf("strings.join %s\n", strings.Join(names, " <> "))

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
	//-- the code was uncommented {end}
}
