package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slicing the Housing Prices
//
//  We have received housing prices. Your task is to print only the requested
//  columns of data (see the expected output).
//
//  1. Separate the data by the newline ("\n") -> rows
//
//  2. Separate each row of the data by the separator (",") -> columns
//
//  3. Find the from and to positions inside the columns depending
//     on the command-line arguments.
//
//  4. Print only the requested column headers and their data
//
//
// RESTRICTIONS
//
//  + You should use slicing when printing the columns.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  go run main.go Location
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  NOTE : Finds the position of the Size column and slices the columns
//         appropriately.
//
//  go run main.go Size
//    Size           Beds           Baths          Price
//
//    100            2              1              100000
//    150            3              2              200000
//    200            4              3              400000
//    500            10             5              1000000
//
//
//  NOTE : Finds the positions of the Size and Baths columns and
//         slices the columns appropriately.
//
//  go run main.go Size Baths
//    Size           Beds           Baths
//
//    100            2              1
//    150            3              2
//    200            4              3
//    500            10             5
//
//
//  go run main.go Beds Price
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  Note : It works even if the given column name does not exist.
//
//  go run main.go Beds NotExist
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  go run main.go NotExist NotExist
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
// Note : It works even if the Price's index > Size's index
//
//        In that case, it resets the invalid starting position to 0.
//
//        But it still uses the Size column.
//
//  go run main.go Price Size
//    Location       Size
//
//    New York       100
//    New York       150
//    Paris          200
//    Istanbul       500
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

// Exercise15 - Slicing the Housing Prices
func Exercise15(vals string) {
	fmt.Println("")
	fmt.Println("---- 16-slices/15-slicing-housing-prices Exercise: Slicing the Housing Prices ----")

	args := strings.Fields(vals)
	fmt.Println("argument string: ", args, "length: ", len(args))

	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	housing := strings.Split(data, "\n")
	headers := strings.Split(housing[0], separator)

	start, end := 0, len(headers)

	// handling of the command line arguments
	for i, v := range headers {
		span := len(args)
		fmt.Printf("for i = %d, v = %s\n", i, v)
		// update start index if the first provided argument
		// matches any value from the column headers
		if span >= 1 && v == args[0] {
			start = i // index
		}

		// update the end position if the second argument
		// provided matches any of the column headers
		if span == 2 && v == args[1] {
			end = i + 1 // index + 1
		}
	}

	// As per Note : It should work even if the Price's index > Size's index
	//        In that case, it resets the invalid starting position to 0.
	//        But it still uses the Size column.
	if start > end {
		start = 0
	}

	fmt.Printf("start: %d | end: %d\n", start, end)

	// for _, v := range headers {
	// 	fmt.Printf("%-15s", v)
	// }
	fmt.Println("\n")

	for i, u := range housing {
		hdata := strings.Split(u, separator)

		for _, h := range hdata[start:end] {
			fmt.Printf("%-15s", h)
		}
		fmt.Println()

		if i == 0 {
			fmt.Println()
		}
	}
}
