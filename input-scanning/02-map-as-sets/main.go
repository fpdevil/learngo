package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// http https://tools.ietf.org/rfc/rfc20.txt | go run main.go network
// curl -s https://tools.ietf.org/rfc/rfc20.txt | go run main.go network

func main() {
	fmt.Println("23-input-scanning/02-map-as-sets")
	fmt.Println()

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please tyoe a search word.")
		return
	}

	query := args[0]

	// create a regular expression matcher to match
	// anything except alphabets to filter the characters
	// like . , ; : etc
	re := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	// testing... prints all words in the text file
	// for in.Scan() {
	// 	fmt.Printf("%s\n", in.Text())
	// }

	// index the words using a map
	var words map[string]bool
	// initialize the map
	words = make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		// using regec matcher here
		word = re.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	// printing the indexed words
	for word := range words {
		fmt.Printf("%s\n", word)
	}

	if words[query] {
		fmt.Printf("input contains the word %q\n", query)
		return
	}

	fmt.Printf("sorry input does not contain the word %q\n", query)

	/*
		// test
		fmt.Println("sun: ", words["sun"], "tesla: ", words["tesla"])

		// query := "sun"
		if _, ok := words[query]; ok {
			fmt.Printf("input contains the word %q\n", query)
			return
		}

		fmt.Printf("sorry input does not contain the word %q\n", query)
	*/
}
