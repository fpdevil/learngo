package main

import "fmt"

// declaring an unused variable at the package level isnot
// a problem and the program will rune
// compiler does not know when or where it willl be used
var unusedPkgLevelVar string

func main() {
	fmt.Println("Variable Declaration")
	fmt.Println("Variable - Name or Identifier - Static Type")

	var nFiles int
	var counter int
	var nCPU int

	fmt.Println("Default Integer Values")
	fmt.Println(
		nFiles,
		counter,
		nCPU,
	)

	var (
		heat float64
		ratio float32
		degree float64
	)
	fmt.Println("Default FloatingPoint Values")
	fmt.Println(
		heat,
		ratio,
		degree,
	)

	var (
		off bool
		valid bool
		closed bool
	)
	fmt.Println("Default Boolean Constants")
	fmt.Println(
		off,
		valid,
		closed,
	)


	var (
		msg string
		name string
		text string
	)
	fmt.Println("Default String Literals")
	fmt.Println(
		msg,
		name,
		text,
	)
	fmt.Println("using Printf for printing an empty string")
	fmt.Printf("(%T) %q, (%T) %q, (%T) %q\n", msg, msg, name, name, text, text)

	var (
		c1 complex64
		c2 complex128
	)
	fmt.Println("Complex types")
	fmt.Println("c1: ", c1, " c2: ", c2)

	var junker rune
	fmt.Println("Rune Types")
	fmt.Println(junker)

	var mybite byte
	fmt.Println("Byte Types")
	fmt.Println(mybite)

	// handling unused variables
	// detection appens during compile time
	var unused string

	// implement a blank identifier to silence the error
	// it eats the values the alive
	// its used mainly for empty returns or skip returns
	_ = unused
}
