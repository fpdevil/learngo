package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

const usage5 = `
	provide two numbers [minval] [maxval] as arguments
`

// Exercise5 - Break Up
func Exercise5(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/05-break-up Exercise: Break Up ----")

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

	var (
		sum    int
		i      = minval
		intstr []string
	)

	for {

		if i > maxval {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}

		sum += i
		if i == maxval {
			intstr = append(intstr, strconv.Itoa(i)+" = ")
		} else {
			intstr = append(intstr, strconv.Itoa(i)+" + ")
		}

		i++
	}

	fmt.Println(strings.Join(intstr, ""), sum)
}
