package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//
//  2. Add book titles to the array
//
//  3. Create two more copies of the array named: upper and lower
//
//  4. Change the book titles to uppercase in the upper array only
//
//  5. Change the book titles to lowercase in the lower array only
//
//  6. Print all the arrays
//
//  7. Observe that the arrays are not connected when they're copied.
//
// NOTE
//  Check out the strings package, it has functions to convert letters to
//  upper and lower cases.
//
// BONUS
//  Invent your own arrays with different types other than string,
//  and do some manipulations on them.
//
// EXPECTED OUTPUT
//   Note: Don't worry about the book titles here, you can use any title.
//
//   books: ["Kafka's Revenge" "Stay Golden" "Everythingship"]
//   upper: ["KAFKA'S REVENGE" "STAY GOLDEN" "EVERYTHINGSHIP"]
//   lower: ["kafka's revenge" "stay golden" "everythingship"]
// ---------------------------------------------------------

// Exercise7 - Assign the Arrays
func Exercise7() {
	fmt.Println("")
	fmt.Println("---- 14 -arrays/07-assign Exercise: Assign the Arrays----")

	books := [...]string{
		"Kafka's Revenge", "Stay Golden", "Everythingship",
	}

	var upper [len(books)]string
	var lower [len(books)]string

	for i, v := range books {
		upper[i] = strings.ToUpper(v)
		lower[i] = strings.ToLower(v)
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}