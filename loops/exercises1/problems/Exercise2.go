package problems

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

// Exercise2 - Sum the Numbers: Verbose Edition
func Exercise2() {
	fmt.Println("")
	fmt.Println("---- 13-loops/02-sum-the-numbers-verbose Exercise: Sum the Numbers: Verbose Edition ----")

	// args := strings.Split(magstr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// if magstr == "" {
	// 	fmt.Println("Give me the magnitude of the earthquake.")
	// 	return
	// }

	const minval, maxval = 1, 10

	var sum int

	for i := minval; i <= maxval; i++ {
		sum += i
		fmt.Print(i)
		if i < maxval {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n", sum)
}
