package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

//  go run main.go "$(cat story.txt)"

func main() {
	fmt.Println("21-project-text-wrapper")
	fmt.Println("Text Wrapper Challenge Guideline")
	fmt.Printf("%s\n", strings.Repeat("-", 40))

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Insufficient Arguments")
		return
	}

	const (
		maxWidth = 40
	)

	var (
		// input text
		input = args[0]

		// track the current line width
		lw int
	)

	// using for range to loop over the runes of a string
	for _, r := range input {
		fmt.Printf("%c", r)

		// if lw > maxWidth {
		// 	// fmt.Printf(" <==[%d]", lw)
		// 	lw = 0
		// 	fmt.Println()
		// } else if r == '\n' {
		// 	lw = 0
		// }
		// lw++

		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}

	}
}
