package problems

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

// Exercise6 - Compare the slices
func Exercise6() {
	fmt.Println("")
	fmt.Println("---- 16-slices/06-compare-the-slices Exercise: Compare the slices ----")

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	var sliceA []string

	sliceA = strings.Split(namesA, ", ")
	// fmt.Println(sliceA)
	sort.Strings(sliceA)
	sort.Strings(namesB)

	// static value to use while print equality message
	var ok string

	// for comparing the slices, we will first check whether
	// the lengths of the slices are same or not
	if len(sliceA) == len(namesB) {
		for i, v := range sliceA {
			// fmt.Println(v, " = ", namesB[i])
			if v != namesB[i] {
				ok = "not "
				break
			}
		}
	}

	fmt.Printf("They are %sequal.\n", ok)

	// alternate approach, but there is not place holder
	// to print the message in case they are not equal
	// if len(sliceA) == len(namesB) {
	// 	for i := range sliceA {
	// 		if sliceA[i] != namesB[i] {
	// 			return
	// 		}
	// 	}
	// 	fmt.Println("They are equal.")
	// }
}
