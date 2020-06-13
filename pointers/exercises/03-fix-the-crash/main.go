// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Fix the crash
//
// EXPECTED OUTPUT
//
//   brand: apple
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	fmt.Println("---- 26-pointers/03-fix-the-crash ----")
	fmt.Println("     EXERCISE: Fix the crash          ")

	// fmt.Printf("main.computer: %T - %#[1]v - %[1]p\n", c)
	var c *computer			// this is a nil pointer

	// populate this with a zero valued struct
	c = &computer{}

	change(c, "apple")
	fmt.Printf("brand: %s\n", *c.brand)
	// fmt.Printf("brand: %s\n", c.brand)
}

func change(c *computer, brand string) {
	// fmt.Printf("computer: %T - %#[1]v - %[1]p\n", c)
	// fmt.Printf("brand: %T\n", brand)
	// the *c pointer is poinitng to a nil pointer causing crash
	// fmt.Printf("%T\n", *c)

	// change below, as brand is a pointer to string
	// (*c.brand) = brand
	c.brand = &brand
}
