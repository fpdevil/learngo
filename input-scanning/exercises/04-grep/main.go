package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------

func main() {
	fmt.Println("--- Exerises 23-input-scanning/04-grep ---")
	fmt.Println("Grep Clone")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println()

	var query string
	if len(os.Args[1:]) != 0 {
		query = os.Args[1]
	}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	for in.Scan() {
		line := in.Text()
		if strings.Contains(line, query) {
			fmt.Printf("%s\n", line)
		}
	}
}
