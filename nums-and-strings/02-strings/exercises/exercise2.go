package exercises

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Print JSON
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

// Exercise2 handles Windows Path
func Exercise2() {
	fmt.Println()
	fmt.Println("---- EXERCISE2: Print JSON from 02-print-json ----")

	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	jsonx := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}
`
	fmt.Println(json)
	fmt.Println("using raw string literals")
	fmt.Println(jsonx)
}
