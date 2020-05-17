package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

// Exercise1 - Declare nil slices
func Exercise1() {
	fmt.Println("")
	fmt.Println("---- 16-slices/01-declare-nil Exercise: Declare nil slices ----")

	// args := strings.Split(magstr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// if magstr == "" {
	// 	fmt.Println("Give me the magnitude of the earthquake.")
	// 	return
	// }

	var (
		names     []string
		distances []int
		data      []byte
		ratio     []float64
		alives    []bool
	)

	fmt.Printf("names:     %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data:      %#v\n", data)
	fmt.Printf("ratio:     %#v\n", ratio)
	fmt.Printf("alives:    %#v\n", alives)

	fmt.Println()
	fmt.Printf("names: %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratio: %T %d %t\n", ratio, len(ratio), ratio == nil)
	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)

	fmt.Println()
	fmt.Printf("names: %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data: %d\n", data)
	fmt.Printf("ratio: %.2f\n", ratio)
	fmt.Printf("alives: %t\n", alives)
}
