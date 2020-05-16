package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

const (
	ops    = "*/+-%"
	noargs = "	Usage: [op=" + ops + "] [size]"
	nosize = `
	Size is missing`
	invalidOp = `
	Invalid operator.
	Valid ops one of: ` + ops
	sizeerr = "Wrong size"
	noOp    = -1
)

// Exercise8 - Math Tables
func Exercise8(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/02-math-tables Exercise: Math Tables ----")

	args := strings.Split(vals, ",")
	fmt.Println("argument string: ", args, "length: ", len(args))

	arglen := len(args)
	switch {
	case vals != "" && arglen == 1:
		fmt.Println(nosize)
		fallthrough
	case vals == "" || arglen < 1:
		fmt.Println(noargs)
		return
	}

	op := args[0]
	max, err := strconv.Atoi(args[1])

	if err != nil || max <= 0 {
		fmt.Println(sizeerr)
		return
	}

	fmt.Println("check: ", strings.IndexAny(ops, op))
	if strings.IndexAny(ops, op) == noOp {
		fmt.Println(invalidOp)
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			var result int

			switch op {
			case "*":
				result = i * j
			case "%":
				if j != 0 {
					result = i % j
				}
			case "/":
				if j != 0 {
					result = i / j
				}
			case "+":
				result = i + j
			case "-":
				result = i - j
			}
			fmt.Printf("%5d", result)
		}
		fmt.Println()
	}

}
