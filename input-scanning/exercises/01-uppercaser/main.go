package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Uppercaser
//
//  Use a scanner to convert the lines to uppercase, and
//  print them.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Print each line.
//
// EXPECTED OUTPUT
//  Please run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	fmt.Println("--- Exerises 23-input-scanning/01-uppercaser ---")
	fmt.Println("Uppercaser")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println()

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	for in.Scan() {
		uc := strings.ToUpper(in.Text())
		fmt.Printf("%v\n", uc)
	}

	if err := in.Err(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
