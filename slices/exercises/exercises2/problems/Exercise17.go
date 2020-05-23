package problems

import (
	"fmt"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort the backing array
//
//  1. Sort only the middle 3 items.
//
//  2. All the slices should see your changes.
//
//
// RESTRICTION
//
//  Do not sort manually. Sort by using the sort package.
//
//
// EXPECTED OUTPUT
//
//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
//
// HINT:
//
//   Middle items are         : [frogger asteroids simcity]
//
//   After sorting they become: [asteroids frogger simcity]
//
// ---------------------------------------------------------

// Exercise17 - Sort the backing array
func Exercise17() {
	fmt.Println("")
	fmt.Println("---- 16-slices/17-internals-backing-array-sort Exercise: Sort the backing array ----")

	// args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	var mid []string
	var i int

	if len(items)%2 != 0 {
		// odd case
		i = (len(items) - 1) / 2
		mid = append(mid, items[i-1:i+2]...)
		// mid = append(mid, items[i-1], items[i], items[i+1])
	} else {
		// even case
		i = len(items) / 2
		mid = append(mid, items[i-1:i+1]...)
		// mid = append(mid, items[i-1], items[i])
	}

	// sort the middle elements
	sort.Strings(mid)

	// fmt.Printf("middle items after: %v\n", mid)
	// fmt.Printf("items %v\n", items[:i-1])
	// append the items starting from the specific index
	items = append(items[:i-1], mid...)

	// extend the slice till it's full capacity
	items = items[:cap(items)]

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
