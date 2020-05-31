package main

import (
	"fmt"
	"os"
)

// input data will be from a flat file which contains the url's
// with http:// prefix

func main() {
	fmt.Println("project-spam-masker/01-step-1")
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
	)

	var (
		input = args[0]
		size  = len(input)

		// create a byte buffer with predefined length for the
		// purpose of output to terminal
		output = make([]byte, size)
	)

	for i := 0; i < size; i++ {
		output = append(output, input[i])

		if len(input[i:]) >= qualen && input[i:i+qualen] == qualifier {
			fmt.Printf("### %s ###\n", string(input[i:i+qualen]))
		}
	}

	fmt.Println("output buffer:")
	fmt.Printf("%v\n", string(output))
}
