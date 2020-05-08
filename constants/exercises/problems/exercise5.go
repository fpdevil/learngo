package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Area
//
//  Fix the following program.
//
// RESTRICTION
//  You should not use any variables.
//
// EXPECTED OUTPUT
//  area = 1250
// ---------------------------------------------------------

// Exercise5 handles Area
func Exercise5() {
	fmt.Println("")
	fmt.Println("---- 10-constants/05-area Exercise: Area ----")

	// const (
	// 	width  int16 = 25
	// 	height int32 = width * 2
	// )
	const (
		width        = 25
		height int32 = width * 2
	)

	// fmt.Printf("area = %d\n", width*height)
	fmt.Printf("area = %d\n", width*height)
}
