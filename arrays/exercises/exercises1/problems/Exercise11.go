package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

// Exercise11 - Find the Average
func Exercise11(vals string) {
	fmt.Println()
	fmt.Println("---- 14-arrays/11-average Exercise: Find the Average ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// handle exactly 5 arguments
	if vals == "" || len(args) < 1 || len(args) > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		numbers [5]float64
		sum     float64
		total   float64
	)

	for i, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		} else {
			numbers[i] = num
			sum += num
			total++
		}
	}

	fmt.Printf("Your numbers: %v\n", numbers)
	fmt.Println("Average:", sum/total)
}
