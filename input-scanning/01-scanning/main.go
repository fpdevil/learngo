package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("23-input-scanning/01-scanning")
	fmt.Println()

	// inducing error
	// os.Stdin.Close()

	// a new inpute reader
	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		fmt.Println("Scanned Text: ", in.Text())
		lines++
	}

	fmt.Printf("There are %d lines in the file\n", lines)

	if err := in.Err(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	// fmt.Printf("Scanned Text: %s\n", in.Text())
	// fmt.Printf("Input Bytes: %d\n", in.Bytes())
}
