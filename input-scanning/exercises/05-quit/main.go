package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {
	fmt.Println("--- Exerises 23-input-scanning/05-quit ---")
	fmt.Println("Quit")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println()

	var words map[string]bool

	in := bufio.NewScanner(os.Stdin)

	words = make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())

		if words[word] {
			fmt.Println("TWICE!")
			return
		}
		words[word] = true
	}
}
