package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
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
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

// Exercise12 - Housing Prices and Averages
func Exercise12() {
	fmt.Println("")
	fmt.Println("---- 16-slices/12-housing-prices-averages Exercise: Housing Prices and Averages ----")

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

	for _, v := range strings.Split(data, "\n") {
		housing := strings.Split(v, ",")
		locations = append(locations, housing[0])

		value, _ := strconv.Atoi(housing[1])
		sizes = append(sizes, value)

		value, _ = strconv.Atoi(housing[2])
		beds = append(beds, value)

		value, _ = strconv.Atoi(housing[3])
		baths = append(baths, value)

		value, _ = strconv.Atoi(housing[4])
		prices = append(prices, value)
	}

	// Rendering the output
	headers := strings.Split(header, separator)
	for i := range headers {
		fmt.Printf("%-15s", headers[i])
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

	fmt.Println(strings.Repeat("=", 75))

	// average values rendering
	// perform the below for matching the output format spacing
	fmt.Printf("%-15s", "")

	// define the size and sum variables (reusable)
	size := len(locations)
	var sum int

	for i := range sizes {
		sum += sizes[i]
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(size))

	sum = 0
	for i := range beds {
		sum += beds[i]
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(size))

	sum = 0
	for i := range baths {
		sum += baths[i]
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(size))

	sum = 0
	for i := range prices {
		sum += prices[i]
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(size))
}
