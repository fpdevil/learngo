package problems

import (
	"bytes"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

// Exercise7 - Append
func Exercise7() {
	fmt.Println("")
	fmt.Println("---- 16-slices/07-append Exercise: Append----")

	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice
	// using the ellipsis operator ...to  unpack the png slice
	header = append(header, png...)

	// 3. compare the slices using the bytes.Equal function
	var ok string
	if !bytes.Equal(png, header) {
		ok = "not "
	}

	// 4. print whether they're equal or not
	fmt.Printf("They are %sequal.\n", ok)
}
