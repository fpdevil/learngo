package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

// Exercise11 - Housing Prices
func Exercise11() {
	fmt.Println("")
	fmt.Println("---- 16-slices/11-housing-prices Exercise: Housing Prices ----")

	// args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// handle arguments
	// if vals == "" || len(args) < 1 {
	// 	fmt.Println("provide a few numbers")
	// 	return
	// }

	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// Load the data into distinct slices
	var (
		locations                  []string
		sizes, beds, baths, prices []int
	)

	headers := strings.Split(header, separator)

	housing := strings.Split(data, "\n")

	for _, v := range housing {
		hdata := strings.Split(v, ",")
		locations = append(locations, hdata[0])

		size, _ := strconv.Atoi(hdata[1])
		sizes = append(sizes, size)

		bed, _ := strconv.Atoi(hdata[2])
		beds = append(beds, bed)

		bath, _ := strconv.Atoi(hdata[3])
		baths = append(baths, bath)

		price, _ := strconv.Atoi(hdata[4])
		prices = append(prices, price)
	}

	for _, v := range headers {
		fmt.Printf("%-15s", v)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 75))

	for i, v := range locations {
		fmt.Printf("%-15s", v)
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d\n", prices[i])
	}
}
