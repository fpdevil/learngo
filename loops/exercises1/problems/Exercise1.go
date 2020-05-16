package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

// Exercise1 - Sum the Numbers
func Exercise1() {
	fmt.Println("")
	fmt.Println("---- 13-loops/01-sum-the-numbers Exercise: Sum the Numbers ----")

	// args := strings.Split(magstr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// if magstr == "" {
	// 	fmt.Println("Give me the magnitude of the earthquake.")
	// 	return
	// }

	var sum int

	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
