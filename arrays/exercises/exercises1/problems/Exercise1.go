package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare empty arrays
//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your best three friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.
//
// HINT
//   When printing the elements of an array, you can use the usual Printf verbs.
//
//   For example:
//     When printing a string array, you can use "%q" verb as usual.
//
// EXPECTED OUTPUT
//  names    : [3]string{"", "", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}
//
//  names    : [3]string
//  distances: [5]int
//  data     : [5]uint8
//  ratios   : [1]float64
//  alives   : [4]bool
//  zero     : [0]uint8
//
//  names    : ["" "" ""]
//  distances: [0 0 0 0 0]
//  data     : [0 0 0 0 0]
//  ratios   : [0.00]
//  alives   : [false false false false]
//  zero     : []
// ---------------------------------------------------------

// Exercise1 - Declare empty arrays
func Exercise1() {
	fmt.Println("")
	fmt.Println("---- 14-arrays/01-declare-empty Exercise: Declare empty arrays ----")

	// args := strings.Split(magstr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// if magstr == "" {
	// 	fmt.Println("Give me the magnitude of the earthquake.")
	// 	return
	// }

	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratio     [1]float64
		alives    [5]bool
		zero      [0]byte
	)

	fmt.Printf("names: %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("ratio: %#v\n", ratio)
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("zero: %#v\n", zero)

	fmt.Println()
	fmt.Printf("names: %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data: %T\n", data)
	fmt.Printf("ratio: %T\n", ratio)
	fmt.Printf("alives: %T\n", alives)
	fmt.Printf("zero: %T\n", zero)

	fmt.Println()
	fmt.Printf("names: %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data: %d\n", data)
	fmt.Printf("ratio: %.2f\n", ratio)
	fmt.Printf("alives: %t\n", alives)
	fmt.Printf("zero: %d\n", zero)
}
