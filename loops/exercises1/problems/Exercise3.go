package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

const usage3 = `
	provide two numbers [minval] [maxval] as arguments
`

// Exercise3 - Sum up to N
func Exercise3(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/03-sum-up-to-n Exercise: Sum up to N ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" {
		fmt.Println(usage3)
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
		sum += i
		if i == maxval {
			intstr = append(intstr, strconv.Itoa(i)+" = ")
		} else {
			intstr = append(intstr, strconv.Itoa(i)+" + ")
		}
	}

	fmt.Println(strings.Join(intstr, ""), sum)
}
