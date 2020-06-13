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
	fmt.Println("---- 26-pointers/05-log-parser ----")
	fmt.Println("EXERCISE: Rewrite the Log Parser program using pointers")

	p := newParser()

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := parse(p, in.Text())
		update(p, &parsed)
	}

	summarize(p)

	// Let's handle the error
	handleErrs([]error{in.Err(), err(p)})
}
