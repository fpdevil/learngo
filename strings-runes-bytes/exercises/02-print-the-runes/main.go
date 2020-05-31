package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	fmt.Println("--- 19-strings-runes-bytes/02-print-the-runes ---")
	fmt.Println("    EXERCISE: Print the runes    ")

	const word = "console"

	//  1. Loop over the "console" word and print its runes one by one,
	//     in decimals, hexadecimals and binary.
	fmt.Printf("%s\n", strings.Repeat("=", 42))
	fmt.Printf("%-06s %-10s %-10s %8s\n", "char", "decimal", "hex", "binary")
	fmt.Printf("%s\n", strings.Repeat("=", 42))

	for i := range word {
		fmt.Printf("%-6c %-10[1]d 0x%-10[1]x 0b%08[1]b\n", word[i])
	}

	//  2. Manually put the runes of the "console" word to a byte slice, one by one.
	//     As the elements of the byte slice use only the rune literals.
	//     Print the byte slice.
	var bs []byte
	for i := range word {
		// fmt.Printf("%c\n", byte(word[i]))
		bs = append(bs, byte(word[i]))
	}
	fmt.Printf("with runes: %s\n", bs)

	//  3. Repeat the step 2 but this time, as the elements of the byte slice,
	//     use only decimal numbers.
	var bss []byte
	for i := range word {
		// fmt.Println(int(word[i]))
		bss = append(bss, byte(int(word[i])))
	}
	fmt.Printf("with decimals: %s\n", bss)

	//  4. Repeat the step 2 but this time, as the elements of the byte slice,
	//     use only hexadecimal numbers.

	for i := range word {
		s := fmt.Sprintf("0x%x", word[i])
		fmt.Printf("%s\n", s)
		// s = fmt.Sprintf("%s0x%x", s, c)
	}

	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))

}
