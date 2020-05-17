package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Assign empty slices
//
//   Assign empty slices to all the slices that you've declared in the previous
//   exercise, and print them here.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 false
//  distances: []int 0 false
//  data     : []uint8 0 false
//  ratios   : []float64 0 false
//  alives   : []bool 0 false
// ---------------------------------------------------------

// Exercise2 - Assign empty slices
func Exercise2() {
	fmt.Println("")
	fmt.Println("---- 16-slices/02-empty Exercise: Assign empty slices ----")

	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	names = []string{}
	distances = []int{}
	data = []byte{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
	// the data is all given in this and it's just a matter of printing
}
