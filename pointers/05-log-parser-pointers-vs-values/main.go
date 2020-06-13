// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("---- 26-pointers/05-log-parser-pointers-vs-values ----")
	fmt.Println()

	p := newParser()                 // Get a new instance of parser
	in := bufio.NewScanner(os.Stdin) // Scan the standard input line by line

	for in.Scan() {
		parsed := parseLines(p, in.Text())
		update(p, parsed)
	}

	// summarize the results
	summarize(p)
	// Let's handle the error
	dumpErrs([]error{in.Err(), err(p)})
}
