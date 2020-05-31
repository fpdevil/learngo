// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("---- 22-maps/01-warm-up ----")

	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	var phones map[string]string

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	var products map[int]bool

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	var mulphones map[string][]string

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	var basket map[int]map[int]int

	fmt.Printf("phone numbers map			: 	%#v\n", phones)
	fmt.Printf("product availability map	: 	%#v\n", products)
	fmt.Printf("multiple phone numbers map	:	%#v\n", mulphones)
	fmt.Printf("shopping basket map			:	%#v\n", basket)
}
