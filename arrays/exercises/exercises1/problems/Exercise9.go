package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

// Exercise9 - Currency Converter
func Exercise9(vals string) {
	fmt.Println()
	fmt.Println("---- 14-arrays/09-currency-converter Exercise: Currency Converter ----")

	args := strings.Split(vals, ",")
	fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	usd, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Printf("%q Invalid amount. It should be a number.", args[0])
		return
	}

	const (
		EUR = iota
		GBP
		JPY
	)

	// using an array with keyed elements and the contants
	// that were declared before
	ratios := [...]float64{
		EUR: 0.92,
		GBP: 0.81,
		JPY: 107.20,
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", usd, ratios[EUR]*usd)
	fmt.Printf("%.2f USD is %.2f GBP\n", usd, ratios[GBP]*usd)
	fmt.Printf("%.2f USD is %.2f JPY\n", usd, ratios[JPY]*usd)
}
