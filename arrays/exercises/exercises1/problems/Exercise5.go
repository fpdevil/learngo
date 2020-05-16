package problems

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------

// Exercise5 - Fix
func Exercise5() {
	fmt.Println("")
	fmt.Println("---- 14-arrays/05-fix Exercise: Fix ----")

	var names [3]string
	names = [3]string{
		"Einstein", "Shepard",
		"Tesla",
	}

	var books [5]string
	books = [5]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everything Together",
		"Star Wars",
		"Harry Potter",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
