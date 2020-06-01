package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

/*
Reference:  https://golang.org/pkg/regexp/syntax/

			GODOCC_STYLE="solarized-dark" godocc -all regexp/syntax
*/

func main() {
	fmt.Println("--- Exerises 23-input-scanning/03-unique-words-2 ---")
	fmt.Println("Unique Words")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println()

	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters

	var (
		words map[string]int

		count int
	)

	// inorder to remove the punctuation characters and
	// numbers, define a regex match for the same
	// the below matches all except mixed alphabets
	re := regexp.MustCompile(`[^A-Za-z]+`)
	// fmt.Printf("using regex %s\n", re.String())

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words = make(map[string]int)
	for in.Scan() {
		count++
		// as per the question
		//  Now the shakespeare.txt contains upper and lower
		//  case letters too.
		// so, to check for uniqueness, we essentially have to
		// convert the words to lowercase
		word := strings.ToLower(in.Text())
		word = re.ReplaceAllString(word, "")
		// fmt.Println(word)
		words[word]++
	}

	fmt.Printf("There are %d words, %d of them are unique.\n", count, len(words))
}
