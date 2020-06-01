package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
// ---------------------------------------------------------

func main() {
	fmt.Println("--- Exerises 23-input-scanning/02-unique-words ---")
	fmt.Println("Unique Words")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println()

	var (
		words map[string]int

		count int
	)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words = make(map[string]int)
	for in.Scan() {
		count++
		words[in.Text()]++
	}
	// fmt.Printf("words map: %#v\n", words)
	fmt.Printf("There are %d words, %d of them are unique.\n", count, len(words))
}
