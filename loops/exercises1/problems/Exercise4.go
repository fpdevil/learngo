package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

const usage4 = `
	provide two numbers [minval] [maxval] as arguments
`

// Exercise4 - Only Evens
func Exercise4(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/04-only-evens Exercise: Only Evens ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" {
		fmt.Println(usage4)
		return
	}

	minval, err := strconv.Atoi(args[0])
	maxval, err := strconv.Atoi(args[1])

	if err != nil || minval > maxval {
		fmt.Printf("invalid arguments [%s, %s] provided\n", args[0], args[1])
		return
	}

	var sum int
	var intstr []string
	for i := minval; i <= maxval; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		if i == maxval {
			intstr = append(intstr, strconv.Itoa(i)+" = ")
		} else {
			intstr = append(intstr, strconv.Itoa(i)+" + ")
		}
	}

	fmt.Println(strings.Join(intstr, ""), sum)
}
