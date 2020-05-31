package main

import (
	"fmt"
	"os"
)

// input data will be from a flat file which contains the url's
// with http:// prefix

func main() {
	fmt.Println("project-spam-masker/03-step-2-no-append")
	fmt.Println("----------------------------\n")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("insufficient arguments")
		return
	}

	// http:// qualifier may be considered as a constant value here
	// so as it's length; so define the same as constants
	// if input data = <text> +  http://xxxxxx  + <http>
	//                 <----> + <-----ql -----> + <---->
	//
	const (
		qualifier = "http://"
		qualen    = len(qualifier)

		// a character to mask the characters of the link
		maskval = '*'
	)

	var (
		input = args[0]
		size  = len(input)

		// create a byte buffer with predefined length for the
		// purpose of output to terminal
		// output = make([]byte, size)
		// output = make([]byte, 0, size)
		output = []byte(input) // create a byte buffer with input string data

		// a boolean variable to find whether the current byte is
		// a part of the url/link or not
		islink bool
	)

	// fmt.Println(string(output))
	// return
	// s.PrintBacking = true
	// s.Show("output", output)

	for i := 0; i < size; i++ {

		if len(input[i:]) >= qualen && input[i:i+qualen] == qualifier {

			// set islink to true indicating that we are inside a link
			islink = true
			// fmt.Printf("### %s ###\n", string(input[i:i+qualen]))

			// add the qualifier value to the output buffer and skip for that
			// much length
			// this is not needed not as the output buffer has full text
			// output = append(output, qualifier...)

			// to skip http:// jump to the index after the length of qualifier
			i += qualen
		}

		// capture the current byte value from input
		c := input[i]

		// in order to decide whether to perform masking of the next few bytes
		// following the link qualifier http://, we first need to decide whether
		// the portion following http:// is a part of link or not
		// we can decide if its a continuous text forming a prt of the link or
		// not by checking if the current byte is one of the below.
		// either a space, newline or a tab character
		switch c {
		case ' ', '\n', '\t':
			islink = false
		}

		// if the current byte is a part of the link, then overwrite it's byte
		// value with the masked character
		if islink {
			// c = maskval
			output[i] = maskval
		}

		// output = append(output, c)
	}

	// s.Show("output", output)
	fmt.Println("output buffer:\n")
	fmt.Printf("%v\n", string(output))
}
