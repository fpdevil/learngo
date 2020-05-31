package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go run main.go 32 127
// go run main.go 9984 10975
// print emojis
// go run main.go 128512 128591

/*
	EXAMPLE UNICODE BLOCKS FROM THE ONLINE TUTORIAL
	------------------------------------------------------------

	1 byte for UTF-8 Encoded
	------------------------------------------------------------
	asciiStart     = '\u0001'      ->  32
	asciiStop      = '\u007f'      ->  127
	upperCaseStart = '\u0041'      ->  65
	upperCaseStop  = '\u005a'      ->  90
	lowerCaseStart = '\u0061'      ->  97
	lowerCaseStop  = '\u007a'      ->  122

	2 bytes for UTF-8 Encoded
	------------------------------------------------------------
	latin1Start    = '\u0080'      ->  161
	latin1Stop     = '\u00ff'      ->  255

	3 bytes for UTF-8 Encoded
	------------------------------------------------------------
	dingbatStart   = '\u2700'      ->  9984
	dingbatStop    = '\u27bf'      ->  10175

	4 bytes for UTF-8 Encoded
	------------------------------------------------------------
	emojiStart     = '\U0001f600'  ->  128512
	emojiStop      = '\U0001f64f'  ->  128591
*/

func main() {
	fmt.Println("bytes-runes-strings-charset-table")

	// var char byte = '⩘'
	// var char rune = '⩘'
	// start, stop := 'A', 'Z'
	// fmt.Println("Unicode code points:", start, stop)
	var (
		start, stop int
	)

	fmt.Println("code point table from", start, "to", stop)
	fmt.Println()

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	// check if the code points are incorrect and reset them
	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "decimal", "hex", "utf-8 enc", strings.Repeat("-", 45))

	// with space between % -12x printf will print
	// bytes separated by spaces
	// string(c) encodes the given code point to a utf-8
	// encoded string value automatically

	for c := start; c <= stop; c++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", c, string(c))
	}
}
