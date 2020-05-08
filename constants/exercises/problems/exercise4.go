package problems

import "fmt"

// ---------------------------------------------------------
// EXERCISE: TAU
//
//  Fix the following program and print the TAU number.
//
// HINT
//  You can use %g verb for printing tau.
//
// EXPECTED OUTPUT
//  tau = 6.283185307179586
// ---------------------------------------------------------

// Exercise4 handles TAU
func Exercise4() {
	fmt.Println("")
	fmt.Println("---- 10-constants/04-tau Exercise: TAU ----")

	// What's the problem with this code?
	// Why it doesn't work?

	// type of tau is calculated from initialized pi
	// it would say pi is undefined with below declaration
	// const pi, tau = 3.14159265358979323846264, pi * 2

	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)
	fmt.Printf("TAU: %g", tau)
}
