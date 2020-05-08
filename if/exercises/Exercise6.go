package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

// Exercise6 - Odd or Even
func Exercise6(agestr string) {
	fmt.Println("")
	fmt.Println("---- 11-if/06-odd-even Exercise: Odd or Even ----")

	args := strings.Split(agestr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if agestr == "" {
		fmt.Println("Pick a number")
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a number\n", args[0])
		return
	}

	if num%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", num)
	} else if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}
}
